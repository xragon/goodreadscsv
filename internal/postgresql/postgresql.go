package postgresql

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/xragon/goodreadscsv/internal/goodreads"

	// "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

// type Book struct {
// 	ID        uuid.UUID    `db:"id"`
// 	Title     string       `db:"title"`
// 	Author    string       `db:"author"`
// 	Rating    int32        `db:"rating"`
// 	DateRead  sql.NullTime `db:"date_read"`
// 	DateAdded sql.NullTime `db:"date_added"`
// 	ISBN      string       `db:"isbn"`
// 	ISBN13    string       `db:"isbn13"`
// 	Status    string       `db:"status"`
// }

// func Read() {
// 	conn, err := pgx.Connect(context.Background(), "postgresql://localhost:5432/books?user=books&password=books")
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// 		os.Exit(1)
// 	}
// 	defer conn.Close(context.Background())
// 	println("connnection successful")

// 	book := Book{}
// 	// err = conn.QueryRow(context.Background(), "SELECT * FROM books").Scan(&book)
// 	// // book, err = conn.Query(context.Background(), "select * from books where")
// 	// if err != nil {
// 	// 	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
// 	// 	os.Exit(1)
// 	// }

// 	row := conn.QueryRow(context.Background(), `SELECT * FROM books`)
// 	switch err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Rating, &book.DateRead, &book.DateAdded, &book.ISBN, &book.ISBN13, &book.Status); err {
// 	// switch err := row.Scan(&book); err {
// 	case sql.ErrNoRows:
// 		fmt.Println("No rows were returned!")
// 	case nil:
// 		fmt.Println(book)
// 	default:
// 		panic(err)
// 	}
// }

func ReadSqlx() {
	db, err := sqlx.Connect("pgx", "postgresql://localhost:5432/books?user=books&password=books")
	if err != nil {
		fmt.Println(err)
	}

	bookid, _ := uuid.NewV4()

	book := goodreads.Book{
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

func Insert() {
	db, err := sqlx.Connect("pgx", "postgresql://localhost:5432/books?user=books&password=books")
	if err != nil {
		fmt.Println(err)
	}

	query := `INSERT INTO books (id, title, author, rating, isbn, isbn13, status) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	book := goodreads.Book{}

	_, err = db.Exec(query, book.ID, book.Title, book.Author, book.Rating, book.ISBN, book.ISBN13, book.Status)
	if err != nil {
		fmt.Println(err)
	}

}
