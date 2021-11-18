package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 使用现有的基础请求对象解析查询字符串参数
	// 示例：URL：/welcome?firstname=Lee&lastname=Moonge
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Lee")
		lastname := c.Query("lastname")  // c.Request.URL.Query().Get("lastname") 的一种快捷方式

		c.String(http.StatusOK, fmt.Sprintf("Hello %v %v", firstname, lastname))
	})
	r.Run()
}
