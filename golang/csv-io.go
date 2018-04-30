package main

import (
	"fmt"
	"os"
	"encoding/csv"
)

type record struct {
	first   string
	last    string
	cell    string
	address string
	salary  string
}

func populateRecords(f1 [][]string, f2 [][]string) []record {
	var r record
	var records []record
	for _, word4reader1 := range f1{
		for _, word4reader2 := range f2{
			if word4reader1[1] == word4reader2[2]{
				r.first = word4reader1[0]
				r.last = word4reader2[0]
				r.cell = word4reader1[1]
				r.address = word4reader1[2]
				r.salary = word4reader2[1]
				records = append(records, r)
			}
		}
	}
	return records
	}

func main() {
	if len(os.Args) < 3 || len(os.Args) > 3 {
		fmt.Println("Please provide 2 filename")
		os.Exit(1)
	}
	file1, file2 := os.Args[1], os.Args[2]
	data4mFile1, err := os.Open(file1)
	data4mFile2, err := os.Open(file2)
	if err != nil {
		panic("Couldn't read the files")
	}
    r1 := csv.NewReader(data4mFile1)
	r2 := csv.NewReader(data4mFile2)
	buffer1, err := r1.ReadAll()
	buffer2, err := r2.ReadAll()
	dump := populateRecords(buffer1, buffer2)
	for _, data := range dump{
		fmt.Println(data)
	}
}


