package main

import (
	"flag"
	"fmt"

	"github.com/xragon/goodreadscsv/internal/goodreads"
)

func main() {
	path := flag.String("path", "goodreads_library_export.csv", "Path and filename to goodreads csv file")
	err := goodreads.Import(*path)
	if err != nil {
		fmt.Println(err)
	}
}
