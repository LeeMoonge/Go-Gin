package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main()  {
	r := gin.Default()

	// 为multipart forms 设置较低的内存限制 (默认是32Mib)
	r.MaxMultipartMemory = 8 << 20  // 8Mib
	r.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// 上传文件至指定目录
		c.SaveUploadedFile(file, "./")

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	r.Run(":8080")
}
