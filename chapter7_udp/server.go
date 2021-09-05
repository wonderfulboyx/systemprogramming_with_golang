package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Server is running at localhost:8888")

	// net.ListenPacket()はnet.Listen()のようにクライアントを待たない
	conn, err := net.ListenPacket("udp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 期待される最大サイズのバッファを用意するか、ヘッダを先読みしてデータ長を確保するといった実装になる
	expectedMaxSize := 1500
	buffer := make([]byte, expectedMaxSize)
	for {
		// conn.Read()は通信内容しか取得できない。
		// UDPはリクエストがあってはじめて接続さきの内容がわかる。
		// 相手に通信を送り返す必要がある場合は、このタイミングにならないと相手のことがわからない
		length, remoteAddress, err := conn.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Recieved: %v: %v\n", remoteAddress, string(buffer[:length]))
		_, err = conn.WriteTo([]byte("Hello from Server"), remoteAddress)
		if err != nil {
			panic(err)
		}
	}
}
