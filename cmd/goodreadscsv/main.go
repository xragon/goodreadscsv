package main

import (
	"fmt"

	"github.com/xragon/goodreadscsv/internal/goodreads"
)

func main() {
	err := goodreads.Import("goodreads_library_export.csv")
	if err != nil {
		fmt.Println(err)
	}
}
