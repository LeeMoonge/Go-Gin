package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
	<title>Https Test</title>
	<script src="/assets/app.js"></script>
</head>
<body>
	<h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func Pusher(c *gin.Context)  {
	if pusher := c.Writer.Pusher(); pusher != nil {
		//使用pusher.Push()做服务器推送
		if err := pusher.Push("/assets/app.js", nil); err != nil {
			log.Printf("Failed to push: %v", err)
		}
	}
	c.HTML(http.StatusOK, "https", gin.H{
		"status": "success",
	})
}

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.SetHTMLTemplate(html)

	r.GET("/", Pusher)

	//监听并在https://127.0.0.1:8080 上启动服务
	_ = r.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}