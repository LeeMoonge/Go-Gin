package main

import "github.com/gin-gonic/gin"

func main()  {
	r := gin.Default()

	r.GET("/asciijson", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "Go",
			"tag": "<br>",
		}
		c.AsciiJSON(200, data)
	})

	r.Run(":8080")
}
