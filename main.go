package main

import (
	"github.com/YoshiRussell/bookclubapp/server/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run()
}
