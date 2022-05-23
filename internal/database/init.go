package database

import (
	"fmt"
	"golayout/pkg/daemon"
	"golayout/pkg/logger"

	"github.com/jmoiron/sqlx"
)

var (
	db     *sqlx.DB
	isInit = false
)

func InitDB(dbOption daemon.DatabaseOption) error {
	if isInit {
		return nil
	}
	isInit = true

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbOption.User, dbOption.Password, dbOption.Host,
		dbOption.Port, dbOption.Name)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	return nil
}

func namedExecX(sqlstr string, arg interface{}) error {
	tx := db.MustBegin()
	_, err := tx.NamedExec(sqlstr, arg)
	if err != nil {
		logger.Error(err)
		return tx.Rollback()
	}
	return tx.Commit()
}

func TableCount(tableName string) uint64 {
	sql := fmt.Sprintf("select count(*) from %s", tableName)
	var result uint64
	err := db.QueryRowx(sql).Scan(&result)
	if err != nil {
		return 0
	}
	return result
}
