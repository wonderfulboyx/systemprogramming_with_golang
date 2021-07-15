package main

import (
	"bytes"
	"fmt"
)

func BufferMain() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())
}
