package database

import (
	"database/sql"
	"github.com/YoshiRussell/bookclubapp/server/models"
)

type Bookstore interface {
	GetALLBooks() ([]models.Book, error)
	Close()
}

// wrapper for database environment
// allows us to decide which database to use
type DbENV struct {
	DB *sql.DB
}

func DatabaseENVInit(local bool, mock bool) (Bookstore, error) {
	if mock {
		return MockDatabaseENVInit()
	}
	
	bookstore := DbENV{}
	var err error
	if local {
		bookstore.DB, err = sql.Open("postgres", "postgres://yoshitest:password@localhost/bookclubtest?sslmode=disable")
	} else {
		bookstore.DB, err = sql.Open("postgres", "someOtherPathForLegitDB")
	}
	return &bookstore, err
}

func (this *DbENV) GetALLBooks() ([]models.Book, error) {
	rows, err := this.DB.Query("SELECT * FROM books;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]models.Book, 0)
	for rows.Next() {
		bk := models.Book{}
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)	
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}

func (this *DbENV) Close() {
	this.DB.Close()
}

