package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/gofrs/uuid"

	// Postgresql Driver
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

// Book is a subsection of a Goodreads CSV entry
type Book struct {
	ID        uuid.UUID    `db:"id"`
	Title     string       `db:"title"`
	Author    string       `db:"author"`
	Rating    int32        `db:"rating"`
	DateRead  sql.NullTime `db:"date_read"`
	DateAdded sql.NullTime `db:"date_added"`
	ISBN      string       `db:"isbn"`
	ISBN13    string       `db:"isbn13"`
	Status    string       `db:"status"`
}

type store struct {
	DB *sqlx.DB
}

// GoodreadsStore exposes prostgres functions for books db
type GoodreadsStore interface {
	WriteRecord(Book) error
}

// NewStore returns an instance of the GoodreadsStore
func NewStore() (GoodreadsStore, error) {
	s := &store{}
	var err error
	s.DB, err = sqlx.Connect("pgx", "postgresql://localhost:5432/books?user=books&password=books")
	if err != nil {
		return nil, err
	}

	return s, nil
}

// ReadSqlx is a practice function to be replaced
func ReadSqlx() {
	db, err := sqlx.Connect("pgx", "postgresql://localhost:5432/books?user=books&password=books")
	if err != nil {
		fmt.Println(err)
	}

	bookid, _ := uuid.NewV4()

	book := Book{
		ID:     bookid,
		Title:  "Example Book",
		Author: "blah",
		Rating: 5,
		ISBN:   "22222",
		ISBN13: "33333",
		Status: "read",
	}

	err = db.Get(&book, "SELECT * FROM books LIMIT 1")

	fmt.Println(book)
}

// WriteRecord will store a Book object in the database
func (s *store) WriteRecord(book Book) error {
	query := `INSERT INTO books (id, title, author, rating, date_read, date_added, isbn, isbn13, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := s.DB.Exec(query, book.ID, book.Title, book.Author, book.Rating, book.DateRead, book.DateAdded, book.ISBN, book.ISBN13, book.Status)
	if err != nil {
		return err
	}

	return nil
}
