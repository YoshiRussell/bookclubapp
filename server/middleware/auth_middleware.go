package middleware

import (
	"fmt"
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

// ** API ** //
var devAud string = "http://localhost:8080"
var prodAud string = "https://nillbookclub/api"

// ** Client ** //
var devIss string = "http://localhost:3000/"
var prodIss string = "https://nillbookclub/"

var jwtMiddleware *jwtmiddleware.JWTMiddleware 

func init() {
	jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options {
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			
			// ** Check that claims[aud] == our API ** //
			checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(devAud, false)
			if !checkAud {
				return token, errors.New("Invalid audience")
			}

			// ** Check that claims[iss] == client **//
			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(devIss, false)
			if !checkIss {
				return token, errors.New("Invalid issuer")
			}

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
	resp, err := http.Get(devIss + ".well-known/jwks.json")
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

