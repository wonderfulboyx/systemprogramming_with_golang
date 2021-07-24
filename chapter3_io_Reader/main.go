package main

import (
	"strings"
)

func main() {
	WriteZipMain()
	var w strings.Builder
	CopyN(&w, strings.NewReader("hogehogehogehoge"), 3)
	println(w.String())
	StreamMain()
}
