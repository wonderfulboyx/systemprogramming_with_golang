package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/http/httputil"
	"strconv"
)

const ResponseMaxSize = math.MaxInt64

func main() {
	addr := "localhost:8888"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer func() {
		println("fuga")
		conn.Close()
	}()

	request, err := http.NewRequest(
		"POST",
		"http://"+addr,
		nil,
	)
	if err != nil {
		panic(err)
	}

	err = request.Write(conn)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(conn)
	response, err := http.ReadResponse(reader, request)
	if err != nil {
		panic(err)
	}

	dump, err := httputil.DumpResponse(response, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))

	if len(response.TransferEncoding) < 1 ||
		response.TransferEncoding[0] != "chunked" {
		panic("wrong transfer encoding")
	}

	for {
		sizeStr, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		size, err := strconv.ParseInt(string(sizeStr[:len(sizeStr)-2]), 16, 64)
		if size == 0 {
			println("hoge")
			break
		}
		if err != nil {
			panic(err)
		}

		line:= make([]byte, int(size))
		io.ReadFull(reader, line)
		reader.Discard(2)
		fmt.Printf( "  %d bytes: %s\n", size, string(line))
	}
	println("hoge")
}
