package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 绑定JSON
type Login struct {
	User string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	// 绑定JSON{{"user":"manu", "password": "123"}}
	r.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "LeeMoonge" || json.Password != "123" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("Hello %s you are logged in !", json.User)})
	})

	// 绑定XML（
	// <?xml version="1.0" encoding="UTF-8"?>
	//	<root>
	//		<user>manu</user>
	//		<password>123</password>
	//	</root>)
	r.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			return
		}

		if xml.User != "LeeMoonge" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("Hello %s you are logged in !", xml.User)})
	})

	// 绑定HTML表单（user=LeeMoonge&password=123）
	r.POST("/someFORM", func(c *gin.Context) {
		var form Login
		// 根据Content-Type Header推断使用哪个绑定器
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			return
		}

		if form.User != "LeeMoonge" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("Hello %s you are logged in !", form.User)})
	})

	r.Run()
}
