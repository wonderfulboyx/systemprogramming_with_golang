package main

import (
	"fmt"
	"strings"
)

func StringBuilderMain() {
	var builder strings.Builder
	builder.Write([]byte("hogehoge"))
	fmt.Println(builder.String())
}
