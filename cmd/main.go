package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

var filename string

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please enter the filename:")
		_, err := fmt.Scanln(&filename)
		if err != nil {
			log.Println(err)
		}
	} else {
		filename = args[1]
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','

	res, err := restoreCSV(reader)
	if err != nil {
		log.Fatalf("Error occured: %s", err)
	}
	fmt.Println(res)
}
