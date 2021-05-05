package routes

import (
	"net/http"
	"fmt"
	"github.com/YoshiRussell/bookclubapp/server/middleware"
	"github.com/YoshiRussell/bookclubapp/server/database"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func SetupRouter(bookStoreDB database.Bookstore) *gin.Engine {
	r := gin.Default()
	
	r.Use(cors.New(cors.Config{
		AllowOrigins:	[]string{"*"},
		AllowMethods: 	[]string{"PUT", "GET", "DELETE", "POST"},
		AllowHeaders:	[]string{"Authorization"},
	}))
	
	r.Use(middleware.AttachDatabase(bookStoreDB))

	r.GET("/", rootHandler)
	r.GET("/mydashboard", middleware.Auth0Middleware(), dashboardHandler)
	r.POST("/mydashboard", middleware.Auth0Middleware(), dashboardHandler2)
	r.GET("/epic", epicHandler)
	return r
}

// Serves index.html containing react script
func rootHandler(c *gin.Context) {
	c.String(http.StatusOK, "This is a placeholder for the homepage.")
}

// This should only be reached by authenticated users
// Must include Authorization header with access token in request
func dashboardHandler(c *gin.Context) {
	useridFromContext, useridExists := c.Get("userid")
	if !useridExists {
		panic("Something went wrong retrieving the userid!")
	}

	Db := c.MustGet("DB").(database.Bookstore)
	userid := fmt.Sprintf("%v", useridFromContext)
	Db.CreateUserIfNew(userid)
	bks, err := Db.GetUsersBooks(userid)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H {
		"username" : "Yoshi",
		"pageNumber" : 22,
		"books" : bks,
	})
}

func dashboardHandler2(c *gin.Context) {
	useridFromContext, useridExists := c.Get("userid")
	if !useridExists {
		panic("Something went wrong retrieving the userid!")
	}

	q := c.Request.URL.Query()
	isbn, ok := q["isbn"]; 
	if !ok {
		panic("isbn not included in query")
	}
	
	Db := c.MustGet("DB").(database.Bookstore)
	userid := fmt.Sprintf("%v", useridFromContext)

	Db.AddBookToUsersBooks(userid, isbn[0])
	c.Status(http.StatusOK)
}

func epicHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "fortnite",
	})
}
