package main

import (
	"bufio"
	"encoding/json"
	"os"
)

func JsonMain() {
	buf := bufio.NewWriter(os.Stdout)
	encoder := json.NewEncoder(buf)
	encoder.SetIndent("", "    ")
	_ = encoder.Encode(map[string]string{
		"example": "hogehoge",
		"hello": "world",
	})
	buf.Flush()
}
