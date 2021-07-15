package main

import "os"

func OsCreateMain() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	_, _ = file.Write([]byte("os.File example\n"))
	_ = file.Close()
}
