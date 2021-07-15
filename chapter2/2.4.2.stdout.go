package main

import "os"

func StdoutMain() {
	_, _ = os.Stdout.Write([]byte("os.Stdout example\n"))
}
