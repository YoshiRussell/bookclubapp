package database

import (
	"github.com/YoshiRussell/bookclubapp/server/models"
	"database/sql"
)

type Bookstore interface {
	CreateUserIfNew(userid string)
	CreateBookIfNew(isbn string)
	Close()
}

// ** IMPLEMENTS THE BOOKSTORE INTERFACE ** //
type MockDb struct {
	DB map[string][]models.Book
}

// ** IMPLEMENTS THE BOOKSTORE INTERFACE ** //
type Db struct {
	DB *sql.DB
}