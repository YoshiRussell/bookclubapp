package database

import (
	"github.com/YoshiRussell/bookclubapp/server/models"
)

func (this *Db) GetALLBooks() ([]models.Book, error) {
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

func (this *Db) Close() {
	this.DB.Close()
}

