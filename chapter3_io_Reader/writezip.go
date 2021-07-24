package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func WriteZipMain()  {
	compress(strings.NewReader("hogehoge"))
}

func compress(r io.Reader) {
	file, _ := os.Create("lenna.zip")
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	w, _ := zipWriter.Create("newfile.txt")

	io.Copy(w, r)
}
