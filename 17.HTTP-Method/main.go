package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 禁用控制台颜色
	gin.DisableConsoleColor()

	// 使用默认中间件（logging 和 recovery 中间件）创建 gin 路由
	r := gin.Default()

	r.GET("/someGet", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	})
	r.POST("/somePost", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"method": "POST"})
	})
	r.PUT("/somePut", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"method": "PUT"})
	})
	r.DELETE("/someDELETE", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"method": "DELETE"})
	})
	r.PATCH("/somePATCH", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"method": "PATCH"})
	})
	r.HEAD("/someHEAD", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"method": "HEAD"})
	})
	r.OPTIONS("/someOPTIONS", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"method": "OPTIONS"})
	})

	_ = r.Run(":8080")
}
