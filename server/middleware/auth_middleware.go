package middleware

import (
	"fmt"
	"log"
	"strings"
	"net/http"
	"errors"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
)

func DummyMiddleware(c *gin.Context) {
	fmt.Println("testing dummy middleware")
	c.Next()
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Use string `json:"use"`
	N   string `json:"n"`
	E   string `json:"e"`
	X5c []string `json:"x5c"`
}

// ** Auth0 API identifier** //
var aud string = "https://nillbookclub/api"

// ** Auth0 tenant identifier ** //
var iss string = "https://dev-35574pmo.us.auth0.com/"

var jwtMiddleware *jwtmiddleware.JWTMiddleware 

func init() {
	jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options {
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			
			// // ** Check that claims[aud] == our API ** //
			// checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
			// if !checkAud {
			// 	return token, errors.New("Invalid audience!")
			// }

			// // ** Check that claims[iss] == client **//
			// checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
			// if !checkIss {
			// 	return token, errors.New("Invalid issuer")
			// }

			// ** Acquire public key from Auth0 tenant endpoint ** //
			cert, err := getPemCert(token)
			if err != nil {
				panic(err.Error())
			}

			// ** Decrypt HASH ** //
			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			return result, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
	})
}

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""

	// ** request for public key from auth0 tenant endpoint ** //
	resp, err := http.Get(iss + ".well-known/jwks.json")
	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)
	if err != nil {
		return cert, err
	}

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATION-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATION-----"
		}
	}

	if cert == "" {
		err := errors.New("Unable to find appropriate key")
		return cert, err
	}

	return cert, nil
}

func Auth0Middleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := jwtMiddleware.CheckJWT(context.Writer, context.Request)
		if err != nil {
			log.Printf("error validating auth token: %v\n", err)
			context.Abort()
			context.Writer.WriteHeader(http.StatusUnauthorized)
			context.Writer.Write([]byte("Unauthorized"))
			return
		}

		authToken := context.Request.Header["Authorization"][0]
		authToken = strings.Split(authToken, "Bearer ")[1]
		parsedToken, _ := jwt.ParseWithClaims(authToken, &jwt.StandardClaims{}, nil)
		if err != nil {
			log.Printf("error parsing token: %v\n", err)
			context.Abort()
			context.Writer.WriteHeader(http.StatusUnauthorized)
			context.Writer.Write([]byte("Unauthorized"))
			return
		}

		tokenData := parsedToken.Claims.(*jwt.StandardClaims)
		log.Printf("claims retrieved %+v\n", tokenData)
		context.Set("userid", tokenData.Subject)
	}
}