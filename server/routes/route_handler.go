package routes

import (
	"net/http"
	"github.com/YoshiRussell/bookclubapp/server/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", rootHandler)
	r.GET("/mydashboard", middleware.DummyMiddleware, dashboardHandler)
	r.GET("/epic", epicHandler)
	return r
}

// Serves index.html containing react script
func rootHandler(c *gin.Context) {
	c.String(http.StatusOK, "This is a placeholder for the homepage.")
}

// This should only be reached by authenticated users
func dashboardHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"username" : "Yoshi",
		"pageNumber" : 22,
	})
}

func epicHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "fortnite",
	})
}
