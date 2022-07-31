package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//router
	r := gin.Default()
	r.Static("/", "./webapp/myapp/dist/myapp")
	r.NoRoute(serveDefault)
	r.Run(":8090")
}

func serveDefault(c *gin.Context) {
	c.File("./webapp/myapp/dist/myapp/index.html")
}
