package database

import (
	"github.com/YoshiRussell/bookclubapp/server/models"
	"github.com/YoshiRussell/bookclubapp/util"
	"fmt"
	"database/sql"
	"strings"
)


func (bookstore *Db) CreateUserIfNew(userid string) {
	bookstore.CreateUserBookRelationTableIfNotExist()
	bookstore.DB.Exec(`INSERT INTO users VALUES ($1) ON CONFLICT (user_id) DO NOTHING;`, userid)
}


func (bookstore *Db) CreateUserBookRelationTableIfNotExist() {
	bookstore.CreateUserTableIfNotExist()
	bookstore.CreateBooksTableIfNotExist()
	_, err := bookstore.DB.Exec(`CREATE TABLE IF NOT EXISTS user_books (
							user_id		CHAR(50) REFERENCES users ON DELETE CASCADE,
							isbn		CHAR(13) REFERENCES books ON DELETE CASCADE,
							PRIMARY KEY (user_id, isbn)
					   );`)
	if err != nil {
		panic(err)
	}
}

func (bookstore *Db) CreateBooksTableIfNotExist() {
	_, err := bookstore.DB.Exec(`CREATE TABLE IF NOT EXISTS books (
							isbn				CHAR(13) PRIMARY KEY		NOT NULL,
							author_firstname	CHAR(50)			NOT NULL,
							author_lastname		CHAR(50)			NOT NULL,
							book_title			CHAR(50)			NOT NULL
					   );`)
	if err != nil {
		panic(err)
	}
}

func (bookstore *Db) CreateUserTableIfNotExist() {
	_, err := bookstore.DB.Exec(`CREATE TABLE IF NOT EXISTS users (
							user_id 	CHAR(50) PRIMARY KEY 	NOT NULL
					   );`)
	if err != nil {
		panic(err)
	}
}

func (bookstore *Db) AddBookToUsersBooks(userid string, isbn string) {
	bookExists := bookstore.rowExists(`SELECT isbn FROM books WHERE isbn=$1`, isbn)
	//userBookRelationExists := bookstore.rowExists(`SELECT isbn FROM users u JOIN user_books ub ON u.user_id = ub.user_id JOIN books b ON ub.isbn = b.isbn`)
	
	fmt.Println(userid)
	if !bookExists {
		bookstore.CreateBookIfNew(isbn)
	}
}

func (bookstore *Db) CreateBookIfNew(isbn string) {
	respBody, err := util.GoogleBooksAPI(isbn)
	if err != nil {
		panic(err)
	}

	bookDetails := util.ParseJSON(respBody)
	author := strings.Split(bookDetails.Authors[0], " ")
	authorFirstName := author[0]
	authorLastName := author[1]
	if len(author) > 2 {
		authorLastName = author[2]
	}
	title := bookDetails.Title
	_, err = bookstore.DB.Exec(`INSERT INTO books VALUES ($1, $2, $3, $4)`, 
		isbn, authorFirstName, authorLastName, title)	
	if err != nil {
		panic(err)
	}
}

func (bookstore *Db) rowExists(query string, args ...interface{}) bool {
	var exists bool
	query = fmt.Sprintf("SELECT EXISTS (%s)", query)
	err := bookstore.DB.QueryRow(query, args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return exists
}

func (bookstore *Db) GetUsersBooks(userid string) ([]models.Book, error) {
	rows, err := bookstore.DB.Query(`SELECT * FROM books WHERE isbn IN (SELECT isbn FROM user_books WHERE user_id = $1)`, userid)
	if err != nil {
		return nil, err
	}

	bks := make([]models.Book, 0)
	for rows.Next() {
		bk := models.Book{}
		err := rows.Scan(&bk.Isbn, &bk.AuthorFirstName, &bk.AuthorLastName, &bk.Title)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	rows.Close()
	return bks, nil
}

// func (bookstore *DB) GetAllUsers() ([]string, error) {

// }

func (bookstore *Db) Close() {
	bookstore.DB.Close()
}