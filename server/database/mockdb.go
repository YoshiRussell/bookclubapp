package database

import (
	"fmt"
	"github.com/YoshiRussell/bookclubapp/server/models"
)

func (this *MockDb) CreateUserIfNew(userid string) {
	if _, ok := this.DB[userid]; !ok {
		this.DB[userid] = make([]models.Book, 0)
	}
}

func (this *MockDb) AddBookToUsersBooks(userid string, isbn string) {
	fmt.Printf(userid)
	fmt.Printf(isbn)
}

func (this *MockDb) GetUsersBooks(userid string) ([]models.Book, error) {
	bks:= make([]models.Book, 0)
	return bks, nil
}

func (this *MockDb) CreateBookIfNew(isbn string) {
	fmt.Printf("created mock book")
}

func (this *MockDb) GetALLUsers() ([]string, error) {
	users := make([]string, 0)
	for user, _ := range this.DB {
		users = append(users, user)
		fmt.Println(user)
	}
	return users, nil
}

func (this *MockDb) Close() {
	fmt.Println("closing mock database")
}

