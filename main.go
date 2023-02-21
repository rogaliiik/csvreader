package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','

	res := restoreCSV(reader)

	fmt.Println(res)
}
