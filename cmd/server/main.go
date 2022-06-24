package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	//r := gin.Default()
	//
	//r.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "hello world")
	//})
	//
	//err := r.Run()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	var id uuid.UUID
	fmt.Println(id == uuid.Nil)
}
