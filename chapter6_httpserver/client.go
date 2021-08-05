package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

var sendMessages = []string{
	"kani", "ebi", "fuga", "hoge",
}

func main() {
	addr := "localhost:8888"
	var conn net.Conn
	var err error
	requests := make([]*http.Request, 0, len(sendMessages))

	conn, err = net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for i, sendMessage := range sendMessages {
		isLastMessage := i == len(sendMessages)-1
		request, err := http.NewRequest(http.MethodGet, "http://"+addr+"?message="+sendMessage, nil)
		if isLastMessage {
			request.Header.Add("Connection", "close")
		} else {
			request.Header.Add("Connection", "keep-alive")
		}
		if err != nil {
			panic(err)
		}
		err = request.Write(conn)
		if err != nil {
			panic(err)
		}
		requests = append(requests, request)
	}

	reader := bufio.NewReader(conn)
	for _, request := range requests {
		response, err := http.ReadResponse(reader, request)
		if err != nil {
			panic(err)
		}
		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
	}
}
