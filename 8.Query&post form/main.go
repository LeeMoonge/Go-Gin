package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostHandler(c *gin.Context)  {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.PostForm("message")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
		"page": page,
		"name": name,
		"message": message,
	})
}

func main() {
	r := gin.Default()

	r.POST("/post", PostHandler)

	r.Run(":8080")
}
