package database

import (
	"fmt"
	"github.com/YoshiRussell/bookclubapp/server/models"
)

type mockENV struct {
	DB map[string]models.Book
}


func MockDatabaseENVInit() (Bookstore, error) {
	bookstore := mockENV{
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

func (this *mockENV) GetALLBooks() ([]models.Book, error) {
	bks := make([]models.Book, 0)
	for _, v := range this.DB {
		fmt.Println(v.Author)
	}
	return bks, nil
}

func (this *mockENV) Close() {
	fmt.Println("closing mock database")
}

