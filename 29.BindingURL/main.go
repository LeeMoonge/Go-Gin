package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	ID string `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}

func main()  {
	r := gin.Default()
	r.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"name": person.Name, "uuid": person.ID})
	})
	r.Run()
}
