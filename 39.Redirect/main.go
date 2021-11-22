package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// HTTP重定向。内部外部重定向均支持
	r.GET("/getredirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	// 通过POST方法进行HTTP重定向
	r.POST("/postredirect", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "https://www.baidu.com")
	})

	// 路由重定向，使用HandleContext
	r.GET("test", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})
	r.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Hello": "Golang"})
	})

	r.Run()
}
