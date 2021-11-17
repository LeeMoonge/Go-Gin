package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	r := gin.Default()

	// 路由组使用gin.BasicAuth()中间件
	// gin.Accounts是map[string]string的一种快捷方式
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo": "bar",
		"austin": "1234",
		"lena": "hello2",
		"manu": "4321",
	}))

	// /admin/secrets 端点
	// 触发"localhost：8080/admin/secrets"
	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取用户，它是由BaseAuth中间件设置的
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})
	_ = r.Run(":8080")
}
