package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//使用单一模板
/*
func main() {
	r := gin.Default()
	//r.LoadHTMLGlob("tempaltes/*")
	r.LoadHTMLFiles("templates/index.tmpl")
	r.GET("/tmpindex", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "哈哈哈啊哈",
		})
	})
	r.Run(":8080")
}
*/

func PostIndex(c *gin.Context) {
	c.HTML( http.StatusOK, "posts/index.tmpl", gin.H{
		"title": "post index",
	})
}

func UserIndex(c *gin.Context)  {
	c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
		"title": "user index",
	})
}

func main()  {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")  //**表示全部文件夹 *表示所有文件
	r.GET("/posts/index", PostIndex)
	r.GET("/users/index", UserIndex)


	_ = r.Run(":8080")
}



