package api

import (
	"net/http"
	"text/template"
)

var (
	tmpl = template.Must(template.New("doc").Parse(`
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">
<html lang="en">
<head>
    <meta http-equiv="Content-Type" content="text/html;charset=UTF-8">
    <title>Document</title>
</head>
<body>
    <h1>Document</h1>

    {{range .}}
    <div>
        <h2 >{{.ApiName}}</h2>
        <span>{{.Method}}</span>
		<a href="{{.Path}}">{{.Path}}</a>
        <p>{{.Doc}}</p>
        <p>Input:</p>
        <pre>
            {{.Input}}
        </pre>
        <p>Output:</p>
        <pre>
            {{.Output}}
        </pre>
    </div>
    {{end}}
</body>
</html>
`))
)

func Doc(w http.ResponseWriter, r *http.Request) {
	err := ctrlServer.GenerateDocument(tmpl, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
