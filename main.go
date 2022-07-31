package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//router 
	r := gin.Default()
	r.GET("/", Get)
	r.Run(":8090")
}

func Get(c *gin.Context) {
	c.String(200,"Hello World!")
}
