package database

import (
	"github.com/YoshiRussell/bookclubapp/server/models"
)

func (bookstore *Db) GetALLBooks() ([]models.Book, error) {
	rows, err := bookstore.DB.Query("SELECT * FROM books;")
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

func (bookstore *Db) CreateUserIfNew(userid string) {
	bookstore.CreateUserBookRelationTableIfNotExist()
	bookstore.DB.Exec(`INSERT INTO users VALUES ($1) ON CONFLICT (user_id) DO NOTHING;`, userid)
}

func (bookstore *Db) Close() {
	bookstore.DB.Close()
}

func (bookstore *Db) CreateUserBookRelationTableIfNotExist() {
	bookstore.CreateUserTableIfNotExist()
	bookstore.CreateBooksTableIfNotExist()
	bookstore.DB.Exec(`CREATE TABLE [IF NOT EXISTS] user_books (
							user_id		INT REFERENCES users (user_id),
							isbn		INT REFERENCES books (isbn),
							CONSTRAINT 	user_books_pkey PRIMARY KEY (user_id, isbn)
					   );`)
}

func (bookstore *Db) CreateBooksTableIfNotExist() {
	bookstore.DB.Exec(`CREATE TABLE [IF NOT EXISTS] books (
							isbn				INT PRIMARY KEY		NOT NULL,
							author_firstname	CHAR(50)			NOT NULL,
							author_lastname		CHAR(50)			NOT NULL,
							book_title			CHAR(50)			NOT NULL
					   );`)
}

func (bookstore *Db) CreateUserTableIfNotExist() {
	bookstore.DB.Exec(`CREATE TABLE [IF NOT EXISTS] users (
							user_id 	CHAR(50) PRIMARY KEY 	NOT NULL
					   );`)
}

