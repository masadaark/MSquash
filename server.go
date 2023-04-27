package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/healthcheck", Healthcheck)
	server.Run(":1234")
}

func Healthcheck(c *gin.Context) {
	c.JSON(200, gin.H{"message": "http status is OK"})
}
