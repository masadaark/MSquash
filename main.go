package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello_world", GetHelloWorld)
	r.Run()
	fmt.Println("Web Server Start")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func GetHelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}
