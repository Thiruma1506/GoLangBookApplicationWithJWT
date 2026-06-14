package model

import "time"

type Book struct {
	BookId     string `json:"bookId" bson:"_id,omitempty"`
	BookName   string `bson:"bookName"`
	BookAuthor string `bson:"bookAuthor"`
	BorrowDate time.Time `bson:"borrowDate"`
}