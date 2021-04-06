package middleware

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func DummyMiddleware(c *gin.Context) {
	fmt.Println("testing dummy middleware")
	c.Next()
}




//var jwtMiddleWare *jwtmiddleware.JWTMiddleware

// func Auth0Middleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		err := jwtMiddleWare.CheckJWT(c.Writer, c.Request)
// 		if err != nil {
// 			log.Printf("error validating autho token: %v\n", err)
// 			c.Abort()
// 			c.Writer.WriteHeader(http.StatusUnauthorized)
// 			c.Writer.Write([]byte("Unauthorized"))
// 			return
// 		}
// 		authToken := c.Request.Header["Authorization"][0]
// 		splitToken := strings.Split(authToken, "Bearer ")
// 		authToken = splitToken[1]
// 		parsedToken, _ := jwt.ParseWithClaims(authToken, &jwt.StandardClaims{}, nil)
// 		if err != nil {
// 			log.Printf("error parsing token: %v\n", err)
// 			c.Abort()
// 			c.Writer.WriteHeader(http.StatusUnauthorized)
// 			c.Writer.Write([]byte("unauthorized"))
// 			return
// 		}
// 		tokenData := parsedToken.Claims.(*jwt.StandardClaims)
// 		log.Printf("claims retrieved %+v\n", tokenData)
// 		c.Next()
// 	}
// }
