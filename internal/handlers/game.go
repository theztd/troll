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

// Default template
var GameTemplate = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<meta http-equiv="refresh" content="3"> 
    <title>Status Page</title>
    <style>
		body {font-family: sans-serif; }
		.grid { display: grid; gap: 1rem; grid-template-columns: repeat(auto-fit, minmax(400px, 1fr)); }
        .div-ok { background-color: #c8e6c9; padding: 10px; margin: 10px; }
        .div-fail { background-color: #ffcdd2; padding: 10px; margin: 10px; }
		.box { border-radius: 12px;	box-shadow: 0 2px 6px rgba(0,0,0,0.1);}
		.box:hover { transform: translateY(-4px); }
        pre { white-space: pre-wrap; word-wrap: break-word; }
    </style>
</head>
<body>
	<div class="grid">
    {{range .}}
		<div class="{{if eq .Code 200 }}div-ok{{else}}div-fail{{end}} box">
			<strong>{{.Url}}</strong><br/>
			<hr noshade>
			<pre>{{.Body}}</pre>
		</div>
    {{end}}
	</div>
</body>
</html>
`

func GameUI(c *gin.Context) {
	results := []adapter.FetchResult{}

	for _, url := range BackendUrls {
		results = append(results, adapter.FetchUrl(url))
	}

	t, err := template.New("status").Parse(GameTemplate)
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
