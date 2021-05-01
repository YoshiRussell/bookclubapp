package database

import (
	"github.com/YoshiRussell/bookclubapp/server/models"
	"github.com/YoshiRussell/bookclubapp/util"
	"database/sql"
	"log"
	"fmt"
)


func DatabaseENVInit(local bool, mock bool) (Bookstore, error) {
	if mock {
		return MockDatabaseENVInit()
	}

	bookstore := Db{}
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Println("cannot load configurations")
		return nil, err
	}
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", 
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName, config.SSLMode)
	log.Printf(connStr)

	if local {
		log.Printf("starting up local database")
		bookstore.DB, err = sql.Open("postgres", connStr)
		//bookstore.DB, err = sql.Open("postgres", "postgres://yoshitest:password@localhost/bookclubtest?sslmode=disable")
	} else {
		bookstore.DB, err = sql.Open("postgres", "someOtherPathForLegitDB")
	}
	return &bookstore, err
}

func MockDatabaseENVInit() (Bookstore, error) {
	bookstore := MockDb{
		DB : make(map[string][]models.Book),
	}
	bookstore.DB["testID"] = append(bookstore.DB["testID"], models.Book {
		Isbn: "testIsbn", 
		AuthorFirstName: "testFirstname",
		AuthorLastName: "testLastname",
		Title: "testTitle", 
	})
	return &bookstore, nil
}