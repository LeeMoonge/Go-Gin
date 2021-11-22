package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginHandler(c *gin.Context)  {
	c.String(http.StatusOK, "LoginHandler")
}

func submitEndpoint(c *gin.Context)  {
	c.String(http.StatusOK, "submitEndpoint")
}

func readEndpoint(c *gin.Context)  {
	c.String(http.StatusOK, "readEndpoint")
}

func main() {
	r := gin.Default()

	// 简单的路由组
	v1 := r.Group("/v1")
	{
		v1.GET("/login", LoginHandler)
		v1.GET("/submit", submitEndpoint)
		v1.GET("/read", readEndpoint)
	}

	// 简单的路由组v2
	v2 := r.Group("/v2")
	{
		v2.POST("/login", LoginHandler)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	r.Run()
}
