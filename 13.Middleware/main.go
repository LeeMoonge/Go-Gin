package main

import "github.com/gin-gonic/gin"

func main()  {
	// 使用gin.New()代替
	r := gin.New()

	// Default 使用 Logger和Recovery中间件
	//r := gin.Default()

	r.Run()
}
