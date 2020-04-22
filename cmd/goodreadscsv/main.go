package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type Book struct {
	GRBookID  string
	Title     string
	Author    string
	Rating    string
	DateRead  string
	DateAdded string
}

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

		bookLine := Book{
			GRBookID:  record[0],
			Title:     record[1],
			Author:    record[2],
			Rating:    record[7],
			DateRead:  record[14],
			DateAdded: record[15],
		}
		fmt.Println(bookLine)

	}
}