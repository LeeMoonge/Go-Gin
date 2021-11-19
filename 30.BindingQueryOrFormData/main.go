package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Person struct {
	Name string `form:"name"`
	Address string `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func TestHandler(c *gin.Context)  {
	var person Person
	// 如果是`GET`请求，只使用`Form`绑定引擎(`query`)
	// 如果是`POST`请求，首先检查`content-type`是否为`JSON`或`XML`，然后再使用`Form`（`form-data`）
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}
	c.String(http.StatusOK, "Success")
}

func main() {
	r := gin.Default()
	r.GET("/testing", TestHandler)
	r.Run()
}
