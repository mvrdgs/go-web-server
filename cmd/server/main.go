package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Pong")
	})

	err := r.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
