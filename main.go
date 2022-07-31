package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.StaticFS("/more_static", http.Dir("my-files"))
	router.Run(":8080")
}
