package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func main() {
	r := gin.Default()

	// gin.H 是map[string]interface{} 的一种快捷方式
	r.GET("/somejson", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "OK"})
	})

	r.GET("/morejson", func(c *gin.Context) {
		// 也可以使用一个结构体
		var msg struct{
			Name string `json:"user"`
			Msg string `json:"msg"`
			Code int `json:"code"`
		}
		msg.Name = "malaohu"
		msg.Code = http.StatusOK
		msg.Msg = "OK"
		// 注意：msg.Name在json中变成了user, Msg在json中变成了msg，Code在json中变成了code
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// protobuf 的具体定义写在 testdata/protoexample 文件中
		data := &protoexample.Test{
			Label: &label,
			Reps: reps,
		}
		// 请注意：数据在响应中变为二进制数据
		// 将输出被 protoexample.Text protobuf 序列化了的数据
		c.ProtoBuf(http.StatusOK, data)
	})

	r.Run()
}
