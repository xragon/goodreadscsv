package goodreads

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

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
func Import(filename string) error {
	csvFile, err := os.Open(filename)

	if err != nil {
		return err
	}
	defer csvFile.Close()

	r := csv.NewReader(csvFile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		bookid, err := uuid.NewV4()
		if err != nil {
			return err
		}

		i64, _ := strconv.ParseInt(record[7], 10, 32)
		rating := int32(i64)

		bookLine := Book{
			ID:        bookid,
			Title:     record[1],
			Author:    record[2],
			ISBN:      strings.Trim(record[5], "=\""),
			ISBN13:    strings.Trim(record[6], "=\""),
			Rating:    rating,
			DateRead:  parseDate(record[14]),
			DateAdded: parseDate(record[15]),
			Status:    record[18],
		}

		fmt.Println(bookLine)
	}
	return nil
}

func parseDate(date string) sql.NullTime {
	layout := "2006/01/02" // 2020/01/29
	if len(date) > 0 {
		parsedDate, _ := time.Parse(
			layout,
			date)
		return sql.NullTime{
			Time:  parsedDate,
			Valid: true,
		}
	}
	return sql.NullTime{
		Valid: false,
	}
}
