package database

import (
	"fmt"
	"github.com/YoshiRussell/bookclubapp/server/models"
)

func (this *MockDb) GetALLBooks() ([]models.Book, error) {
	bks := make([]models.Book, 0)
	for _, v := range this.DB {
		fmt.Println(v.Author)
	}
	return bks, nil
}

func (this *MockDb) Close() {
	fmt.Println("closing mock database")
}

