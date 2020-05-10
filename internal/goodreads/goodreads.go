package goodreads

import (
	"database/sql"

	"github.com/gofrs/uuid"
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

// Import a segment section of a Goodreads CSV in postgresql
func Import() {

}
