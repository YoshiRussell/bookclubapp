package main

import (
	"github.com/YoshiRussell/bookclubapp/server/routes"
	"github.com/YoshiRussell/bookclubapp/server/database"
	_ "github.com/lib/pq"
	"fmt"
)

func main() {

	// boolean defines if we are in testing environment
	bookstore, err := database.DatabaseENVInit(true, true)
	if err != nil {
		panic(err)
	}
	defer bookstore.Close()
	
	bks, err := bookstore.GetALLBooks()
	if err != nil {
		panic(err)
	}
	for _, bk := range bks {
		fmt.Printf("%s, %s, %s, $%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}

	router := routes.SetupRouter()
	router.Run()
}

