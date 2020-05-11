package goodreads

import (
	"database/sql"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/xragon/goodreadscsv/internal/postgresql"
)

// Import a segment section of a Goodreads CSV in postgresql
func Import(filename string) error {
	csvFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	gr, err := postgresql.NewStore()
	if err != nil {
		return err
	}

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

		bookLine := postgresql.Book{
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

		// fmt.Println(bookLine)
		gr.WriteRecord(bookLine)
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
