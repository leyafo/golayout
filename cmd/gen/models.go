package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"golayout/pkg/daemon"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"golayout/pkg/strutil"

	"go/format"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/pflag"
)

type columnInfo struct {
	DbTag    string
	Type     string
	Field    string
	Comments string
}

type tableInfo struct {
	TableName  string
	Columns    []columnInfo
	StructName string
	Var        string
	Imports    []string
}

type MysqlDesc struct {
	Field   string         `db:"Field"`
	Type    string         `db:"Type"`
	Null    string         `db:"Null"`
	Key     string         `db:"Key"`
	Default sql.NullString `db:"Default"`
	Extra   string         `db:"Extra"`
}

var (
	configFile string
	modelPath  string
)

func init() {
	pflag.StringVarP(&configFile, "config", "f", "", "load the config file name")
	pflag.StringVarP(&modelPath, "registerAddr", "m", "", "listening the special address")
	pflag.Parse()
}

func main() {
	var apiOpt daemon.ApiOption
	err := daemon.ParseConfig(configFile, &apiOpt)
	if err != nil {
		panic(err.Error())
	}
	daemon.SetGlobalApiOption(&apiOpt)

	db, err := sqlx.Connect("mysql", apiOpt.Database.DBString("mysql"))
	if err != nil {
		panic(err.Error())
	}

	var tableNames []string
	err = db.Select(&tableNames, `show tables`)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(tableNames)

	for _, tableName := range tableNames {
		var ti tableInfo
		ti.TableName = tableName
		ti.StructName = strutil.ToCamelCase(ti.TableName)
		ti.Var = ti.TableName[0:1]
		rows, err := db.Queryx("describe " + ti.TableName)
		if err != nil {
			panic(err)
		}
		for rows.Next() {
			var mysqlDesc MysqlDesc
			err := rows.StructScan(&mysqlDesc)
			if err != nil {
				panic(err)
			}
			var ci columnInfo
			if strutil.InsensitiveCmp(mysqlDesc.Field, "id") ||
				strings.Contains(strings.ToLower(mysqlDesc.Field), "_id") {
				ci.Type = "uint64"
				ci.DbTag = mysqlDesc.Field
			} else {
				if strings.ToLower(mysqlDesc.Type) == "tinyint(1)" {
					mysqlDesc.Type = "bool"
				}
				trimIndex := strings.Index(mysqlDesc.Type, "(")
				if trimIndex == -1 {
					trimIndex = len(mysqlDesc.Type)
				}
				ci.Type = mappingDbType(mysqlDesc.Type[0:trimIndex], mysqlDesc.Null == "YES")
			}
			ci.Field = strutil.ToCamelCase(mysqlDesc.Field)
			ci.DbTag = mysqlDesc.Field
			if mysqlDesc.Default.String == "" {
				mysqlDesc.Default.String = "NULL"
			}
			ci.Comments = fmt.Sprintf("%s NULL:%s Default:%s", mysqlDesc.Field, mysqlDesc.Null, mysqlDesc.Default.String)
			ti.Columns = append(ti.Columns, ci)
			ti.Imports = appendImport(ti.Imports, ci.Type)
		}
		rows.Close()
		fmt.Println("gentreate table:", ti.TableName)
		if err = genCode(modelPath, ti); err != nil {
			panic(err)
		}
	}
}

func appendImport(imports []string, typ string) []string {
	typ = strings.TrimSpace(typ)
	var i string
	switch typ {
	case "decimal.Decimal":
		i = "github.com/shopspring/decimal"
	default:
		i = ""
	}
	if strings.Contains(typ, "sql.") {
		i = "database/sql"
	}
	if i != "" {
		for _, im := range imports {
			if im == i {
				return imports
			}
		}
		imports = append(imports, i)
	}
	return imports
}

func genCode(modelPath string, ti tableInfo) error {
	var (
		err error
		buf = new(bytes.Buffer)
	)
	tmpl := template.New("model.go.tpl").Funcs(template.FuncMap{
		"insertStr":    insertStr,
		"updateStr":    updateStr,
		"stateTypeStr": stateTypeStr,
	})
	tmpl, err = tmpl.ParseFiles(filepath.Join(modelPath, "model.go.tpl"))
	if err != nil {
		return err
	}
	err = tmpl.Execute(buf, ti)
	if err != nil {
		return err
	}

	formatedText, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	f, err := os.OpenFile(path.Join(modelPath, ti.TableName+"_gen.go"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	f.Write(formatedText)
	return nil
}

func insertStr(columns []columnInfo) string {
	var colNames, colValues string
	colNames = "("
	colValues = "("
	notFirst := false
	for _, ci := range columns {
		if ci.DbTag == "id" {
			continue
		}
		if notFirst {
			colNames += ","
			colValues += ","
		}
		colNames += ci.DbTag
		colValues += ":" + ci.DbTag
		notFirst = true
	}
	colNames += ")"
	colValues += ")"
	return colNames + " VALUES " + colValues
}

func updateStr(columns []columnInfo) string {
	var s string
	s = "SET "
	notFirst := false
	for _, ci := range columns {
		if ci.DbTag == "id" {
			continue
		}
		if notFirst {
			s += ","
		}
		s += ci.DbTag + "=:" + ci.DbTag
		notFirst = true
	}
	return s
}

func stateTypeStr(columnInfo []columnInfo, structName string) string {
	var s string
	for i, ci := range columnInfo {
		if strings.Contains(strings.ToLower(ci.Field), "state") {
			s += fmt.Sprintf("type %s%sType string \n", structName, ci.Field)
			ci.Type = fmt.Sprintf("%s%sType", structName, ci.Field)
			columnInfo[i] = ci
		}
	}
	return s
}

func mappingDbType(dbType string, nullAllow bool) string {
	var typ string
	switch strings.ToLower(dbType) {
	case "boolean", "bool":
		typ = "bool"
		if nullAllow {
			typ = "sql.NullBool"
		}
	case "varchar", "char", "binary", "varbinary", "text", "tinytext", "mediumtext", "longtext":
		typ = "string"
		if nullAllow {
			typ = "sql.NullString"
		}
	case "blog", "mediumblob", "longblob", "tinyblob":
		typ = "sql.RawBytes"
	case "tinyint", "smallint", "mediumint", "int", "integer":
		typ = "int"
		if nullAllow {
			typ = "sql.NullInt32"
		}
	case "bigint":
		typ = "int64"
		if nullAllow {
			typ = "sql.NullInt64"
		}
	case "float", "double", "real":
		typ = "float64"
		if nullAllow {
			typ = "sql.NullFloat64"
		}
	case "decimal", "dec":
		typ = "decimal.Decimal"
	case "date", "timestamp", "datetime", "time":
		typ = "time.Time"
		if nullAllow {
			typ = "sql.NullTime"
		}
	case "year":
		typ = "int"
	case "json":
		typ = "sql.RawBytes"
	default:
		panic(fmt.Sprintf("Cannot convert type: %s", dbType))
	}
	return typ
}
