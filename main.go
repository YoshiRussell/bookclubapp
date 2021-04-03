package main

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/epic", epicHandler)
	return r
}
func main() {
	router := SetupRouter()
	router.Run()
}

func epicHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "fortnite",
	})
}
