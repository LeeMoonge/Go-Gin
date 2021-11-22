package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

var g errgroup.Group

func route01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery(), gin.Logger())
	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg": "Welcome server 01",
		})
	})
	return e
}

func route02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery(), gin.Logger())
	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg": "Welcome server 02",
		})
	})
	return e
}

func main() {
	server01 := &http.Server{
		Addr: ":8080",
		Handler: route01(),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr: ":8081",
		Handler: route02(),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
