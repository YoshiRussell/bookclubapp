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
	userid, useridExists := c.Get("userid")
	if !useridExists {
		panic("Something went wrong retrieving the userid!")
	}

	Db := c.MustGet("DB").(database.Bookstore)
	
	bks, err := Db.GetALLBooks()
	if err != nil {
		panic(err)
	}

	for _, bk := range bks {
		fmt.Printf(bk.Author)
	}
	//Db.CreateUserIfNew(userid);
	
	fmt.Println(userid)

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
