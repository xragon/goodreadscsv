package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

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

		fmt.Println(record[0])
		break
	}
}
