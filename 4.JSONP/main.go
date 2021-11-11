package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonP(c *gin.Context) {
	data := map[string]interface{}{
		"foo": "bar",
	}
	// JSONP?callback=x
	// 将输出：x({\"foo\":\"bar\"})
	c.JSONP(http.StatusOK, data)
}

func main()  {
	r := gin.Default()
	r.GET("/JSONP", JsonP)

	_ = r.Run(":8080")
}
