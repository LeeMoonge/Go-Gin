package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SecureJson(c *gin.Context) {
	names := []string{"aab", "bba", "aabb"}

	// 将输出: while(1);["aab", "bba", "aabb"]
	c.SecureJSON(http.StatusOK, names)
}

func main() {
	r := gin.Default()

	// 也可以使用自己的SecureJSON前缀
	// r.SecureJsonPrefix(")]}', \n")

	r.GET("/securejson", SecureJson)

	r.Run(":8080")
}
