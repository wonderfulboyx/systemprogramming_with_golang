# UDPソケットを使ったマルチキャスト通信
## TCPとUDP
- TCPもUDPもトランスポート層プロトコル
- TCPはコネクションの管理、データロス検知、通信制限、パケットの順序などができる
- 一方UDPは一方的にデータを送りつけるだけで上記のTCPができるような高度な管理はできない。
- TCPができないがUDPができることとして、マルチキャストとブロードキャストがある
  - どちらも複数のコンピュータに同時にメッセージを送ることができる
- UDPにはコネクションやハンドシェイクがないため、その分のオーバーヘッドを省略できるので高速といわれる

## 現代のUDP
- ほとんどの場合TCPを選択して良い
- アプリケーション開発の観点で選択するなら次のような特殊な条件に合致する場合だけ
  - ロスしても良い
  - 順序が変わってもアプリケーションレイヤーでカバーできる
  - マルチキャストが必要
  - ハンドシェイクの手間すら惜しい

## サーバー側の実装例
```go
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
```

## POSIXとの違い
POSIXでは`listen()`や`accept()`、`connect()`などをつかわずに直接`recvFrom()`や`sendto()`で通信を行えるが、
Goの場合はTCPと同じような準備が必要になる。この点はPOSIXに詳しい人は違和感を覚える点で、一般的な実装ではない。

## UDPを使った実世界のサンプル

### NTP
https://medium.com/learning-the-go-programming-language/lets-make-an-ntp-client-in-go-287c4b9a969f

NTP(Network Time Protocol)の実装例。NTPは正しい時刻を同期するためのプロトコル

### peerdiscovery
https://github.com/schollz/peerdiscovery

マルチキャストの実装例

## TCPとUDPの機能の違い

### TCPの再送処理とフロー制御

- 再送処理
  - 受信側はシーケンス番号とペイロードサイズを送信側に送り返す
  - 送信側は受信側の応答がない場合、落ちたと判断してもう一度送る
- フロー制御
  - 受信側が受信用バッファのサイズを送信側に伝え、そのサイズまでは受信側の確認を待たずに送信できる仕組み。受信側のサイズはウィンドウサイズとよばれ、コネクション確立時に確認し合う。
  - 受信側で読み込み処理が間に合わない場合は、ウィンドウサイズを送信側に伝えて通信量を制御できる = フロー制御

## UDPではフレームサイズを気にする必要がある
- ひとかたまりで送信できるデータの大きさの上限を最大転送単位(MTU)という。
- MTUに収まらないデータは、TCP/UDPよりも下のレイヤーであるIPレイヤーで複数のパケットに分割される。これをIPフラグメンテーションと呼ぶ。
- 分割されたデータはIPレイヤーで再結合されるが、カーネル内部の結合待ちが発生するためUDPの応答性の高さが損なわれてしまう。
- そのため、フレームサイズを気にしてアプリケーションの設計を行う必要がある。
