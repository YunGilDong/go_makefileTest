package main

import (
	"data"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	var lcdata data.LC
	lcdata.LC_ID = 1
	fmt.Println("Hello, My friend!")
	fmt.Println(lcdata)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
