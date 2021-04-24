package middleware

import (
	"github.com/YoshiRussell/bookclubapp/server/database"
	"github.com/gin-gonic/gin"
)

func AttachDatabase(bookstoreDB database.Bookstore) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("DB", bookstoreDB)
		context.Next()
	}
}