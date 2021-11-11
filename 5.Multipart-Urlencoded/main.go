package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginFrom struct {
	User string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func LoginHandler(c *gin.Context)  {
	//可以使用显式绑定声明绑定 multipart form:
	//c.ShouldBindWith(&form, binding.Form)
	//或者简单的使用ShouldBind方法自动绑定:
	var form LoginFrom
	// 在这种情况下，将自动选择合适的绑定
	if c.ShouldBind(&form) == nil {
		if form.User == "user" && form.Password == "password" {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	}
}

func main() {
	r := gin.Default()
	r.POST("/login", LoginHandler)

	_ = r.Run(":8080")
}
