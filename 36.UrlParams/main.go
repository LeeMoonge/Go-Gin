package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 此handler将匹配/user/john 但不会匹配/user/或者/user
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 此handler将匹配/user/john/和/user/john/send
	// 如果没有其他路由匹配/user/john，他将重定向到/user/john/
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + "is" + action
		c.String(http.StatusOK, message)
	})

	r.Run()
}
