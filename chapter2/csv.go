package main

import (
	"encoding/csv"
	"io"
	"os"
)

func GenerateExampleCsv(writer io.Writer) {
	w := csv.NewWriter(writer)
	w.UseCRLF = true
	data := [][]string{
		{"hoge", "fuga", "piyo"},
		{"1", "3", "2"},
		{"1", "3", "2"},
		{"1", "3", "2"},
		{"1", "3", "2"},
		{"1", "3", "2"},
		{"1", "3", "2"},
		{"1", "3", "2"},
	}
	w.WriteAll(data)
	w.Flush()
}

func CsvMain() {
	file, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	GenerateExampleCsv(file)
}
