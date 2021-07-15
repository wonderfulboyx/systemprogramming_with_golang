package main

import (
	"io"
	"net"
	"os"
)

func NetDialMain() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	_, _ = io.WriteString(conn, "ほげほげ")
	io.Copy(os.Stdout, conn)
}
