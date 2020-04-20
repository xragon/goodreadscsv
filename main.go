package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

type Book struct {
	Title  string
	Author string
}

func main() {
	dat, err := ioutil.ReadFile("goodreads_library_export.csv")
	if err != nil {
		panic(err)
	}
	// r := csv.NewReader(dat)
	r := csv.NewReader(strings.NewReader(string(dat)))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if record[18] == "to-read" {
			bookLine := Book{
				Title:  record[1],
				Author: record[2],
			}
			fmt.Println(bookLine.Title + ", " + bookLine.Author)
		}

	}
}
