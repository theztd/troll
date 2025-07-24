package handlers

import (
	"bytes"
	"html/template"

	"github.com/gin-gonic/gin"
	"gitlab.com/theztd/troll/internal/adapter"
)

var BackendUrls = []string{
	"http://service-a/v1/info",
	"http://service-b/v1/info",
	"http://service-c/v1/info",
}

const html_template = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>Status Page</title>
    <style>
        .div-ok { background-color: #c8e6c9; padding: 10px; margin: 10px; }
        .div-fail { background-color: #ffcdd2; padding: 10px; margin: 10px; }
        pre { white-space: pre-wrap; word-wrap: break-word; }
    </style>
</head>
<body>
    {{range .}}
    <div class="{{if eq .Code 200 }}div-ok{{else}}div-fail{{end}}">
        <strong>{{.Url}}</strong><br/>
        <pre>{{.Body}}</pre>
    </div>
    {{end}}
</body>
</html>
`

func StatusNice(c *gin.Context) {
	results := []adapter.FetchResult{}

	for _, url := range BackendUrls {
		results = append(results, adapter.FetchUrl(url))
	}

	t, err := template.New("status").Parse(html_template)
	if err != nil {
		c.String(500, "Template error: %v", err)
		return
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, results); err != nil {
		c.String(500, "Render error: %v", err)
		return
	}

	c.Data(200, "text/html; charset=utf-8", buf.Bytes())
}
