package database

import (
	"github.com/YoshiRussell/bookclubapp/server/models"
	"database/sql"
)

func DatabaseENVInit(local bool, mock bool) (Bookstore, error) {
	if mock {
		return MockDatabaseENVInit()
	}

	bookstore := Db{}
	var err error
	if local {
		bookstore.DB, err = sql.Open("postgres", "postgres://yoshitest:password@localhost/bookclubtest?sslmode=disable")
	} else {
		bookstore.DB, err = sql.Open("postgres", "someOtherPathForLegitDB")
	}
	return &bookstore, err
}

func MockDatabaseENVInit() (Bookstore, error) {
	bookstore := MockDb{
		DB : make(map[string]models.Book),
	}
	bookstore.DB["testID"] = models.Book {
		Isbn: "testIsbn", 
		Title: "testTitle", 
		Author: "testAuthor", 
		Price: 69.99,
	}
	return &bookstore, nil
}