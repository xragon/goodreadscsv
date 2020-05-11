package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/xragon/goodreadscsv/internal/goodreads"
)

// type Book struct {
// 	Title     string
// 	Author    string
// 	Rating    string
// 	DateRead  string
// 	DateAdded string
// 	ISBN      string
// 	ISBN13    string
// 	Status    string
// }

// func main() {
// 	// postgresql.Read()
// 	// postgresql.ReadSqlx()
// 	postgresql.Insert()
// }

func main() {
	// dat, err := ioutil.ReadFile("goodreads_library_export.csv")
	// if err != nil {
	// 	panic(err)
	// }
	// r := csv.NewReader(dat)
	csvFile, err := os.Open("goodreads_library_export.csv")

	if err != nil {
		panic(err)
	}
	defer csvFile.Close()
	// r := csv.NewReader(strings.NewReader(string(dat)))
	r := csv.NewReader(csvFile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		i64, _ := strconv.ParseInt(record[7], 10, 32)
		rating := int32(i64)

		// layoutISO := "2006/01/02" // 2020/01/29

		// var dateread time.Time

		// if len(record[14]) > 0 {
		// 	dateread, _ = time.Parse(
		// 		layoutISO,
		// 		record[14])
		// }

		// dateadded, _ := time.Parse(
		// 	layoutISO,
		// 	record[15],
		// )

		// fmt.Println("" + len(record[14]))
		fmt.Printf("Read: %s\n", record[14])
		fmt.Printf("Added: %s\n", record[15])
		// fmt.Println(dateread, dateadded)

		bookLine := goodreads.Book{
			Title:  record[1],
			Author: record[2],
			ISBN:   strings.Trim(record[5], "=\""),
			ISBN13: strings.Trim(record[6], "=\""),
			Rating: rating,
			// DateRead:  dateread,
			// DateAdded: record[15],
			DateRead:  parseDate(record[14]),
			DateAdded: parseDate(record[15]),
			Status:    record[18],
		}

		// bookLine.DateAdded.Time = dateadded
		// bookLine.DateAdded.Valid = true

		// bookLine.DateRead = sql.NullTime{
		// 	Time:  dateread,
		// 	Valid: true,
		// }

		fmt.Println(bookLine)

	}
}

func parseDate(date string) sql.NullTime {
	layoutISO := "2006/01/02" // 2020/01/29
	if len(date) > 0 {
		parsedDate, _ := time.Parse(
			layoutISO,
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
