package main

import (
	"github.com/YoshiRussell/bookclubapp/server/routes"
	"github.com/YoshiRussell/bookclubapp/server/database"
	_ "github.com/lib/pq"
)

func main() {

	// boolean defines if we are in testing environment
	bookstoreDB, err := database.DatabaseENVInit(true, false)
	if err != nil {
		panic(err)
	}
	err = bookstoreDB.Ping()
	if err != nil {
		panic(err)
	}
	defer bookstoreDB.Close()

	router := routes.SetupRouter(bookstoreDB)
	router.Run()
}

