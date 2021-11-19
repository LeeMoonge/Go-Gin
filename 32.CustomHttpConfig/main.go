package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//func main() {
//	router := gin.Default()
//	http.ListenAndServe(":8080", router)
//}

func main() {
	r := gin.Default()

	s := &http.Server{
		Addr: ":8080",
		Handler: r,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
