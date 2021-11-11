package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func FormPostHandler(c *gin.Context)  {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "nil")
	c.JSON(http.StatusOK, gin.H{
		"status": "posted",
		"message": message,
		"nick": nick,
	})

}

func main() {
	r := gin.Default()
	r.POST("/form_post", FormPostHandler)

	_ = r.Run(":8080")
}

//使用 curl -v --form message=hahaha --form nick=aabb http://localhost:8080/form_post 命令测试
