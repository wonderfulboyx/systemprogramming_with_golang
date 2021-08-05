package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

// https://www.aozora.gr.jp/cards/000119/files/43043_17363.html
var contents = []string{
	"ある時はヘーゲルが如萬有をわが體系に統すべんともせし",
	"ある時はアミエルが如つゝましく息をひそめて生きんと思ひし",
	"ある時は若きジイドと諸共に生命に充ちて野をさまよひぬ",
	"ある時はヘルデルリンと翼はね竝べギリシャの空を天翔りけり",
	"ある時はフィリップのごと小ちさき町に小ちひさき人々ひとを愛せむと思ふ",
	"ある時はラムボーと共にアラビヤの熱き砂漠に果てなむ心",
	"ある時はゴッホならねど人の耳を喰ひてちぎりて狂はんとせし",
	"ある時は淵明えんめいが如疑はずかの天命を信ぜんとせし",
	"ある時は觀念イデアの中に永遠を見んと願ひぬプラトンのごと",
	"ある時はノ※(濁点付き片仮名ワ、1-7-82)ーリスのごと石に花に奇しき祕文を讀まむとぞせし",
	"ある時は人を厭ふと石の上に默もだもあらまし達磨の如く",
	"ある時は李白の如く醉ひ醉ひて歌ひて世をば終らむと思ふ",
	"ある時は王維をまねび寂じやくとして幽篁の裏うちにひとりあらなむ",
	"ある時はスウィフトと共にこの地球ほしのYahooヤフー 共をば憎みさげすむ",
	"ある時はヴェルレエヌの如雨の夜の巷に飮みて涙せりけり",
	"ある時は阮籍げんせきがごと白眼に人を睨みて琴を彈ぜむ",
	"ある時はフロイドに行きもろ人の怪あやしき心理こころさぐらむとする",
	"ある時はゴーガンの如逞ましき野生なまのいのちに觸ればやと思ふ",
	"ある時はバイロンが如人の世の掟おきて踏躪り呵々と笑はむ",
	"ある時はワイルドが如深き淵に墮ちて嘆きて懺悔せむ心",
	"ある時はヴィヨンの如く殺あやめ盜み寂しく立ちて風に吹かれなむ",
	"ある時はボードレエルがダンディズム昂然として道行く心",
	"ある時はアナクレオンとピロンのみ語るに足ると思ひたりけり",
	"ある時はパスカルの如心いため弱き蘆をば讚ほめ憐れみき",
	"ある時はカザノ※(濁点付き片仮名ワ、1-7-82)のごとをみな子の肌をさびしく尋とめ行く心",
	"ある時は老子のごとくこれの世の玄のまた玄空しと見つる",
	"ある時はゲエテ仰ぎて吐息しぬ亭々としてあまりに高し",
	"ある時は夕べの鳥と飛び行きて雲のはたてに消えなむ心",
	"ある時はストアの如くわが意志を鍛へんとこそ奮ひ立ちしか",
	"ある時は其角の如く夜の街に小傾城などなぶらん心",
	"ある時は人麿のごと玉藻なすよりにし妹をめぐしと思ふ",
	"ある時はバッハの如く安らけくたゞ藝術に向はむ心",
}

func isGZipAcceptable(request *http.Request) bool {
	return strings.Contains(
		strings.Join(request.Header["Accept-Encoding"], ","), "gzip",
	)
}

func processSession(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	for {
		// タイムアウトを設定
		err := conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		if err != nil {
			panic(err)
		}

		// リクエストの読み込み
		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			neterr, ok := err.(net.Error)
			if ok && neterr.Timeout() {
				fmt.Println("Timeout")
				break
			} else if err == io.EOF {
				break
			}
			panic(err)
		}
		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
		fmt.Fprintf(conn, strings.Join([]string{
			"HTTP/1.1 200 OK",
			"Content-Type: text/plain",
			"Transfer-Encoding: chunked",
			"",
			"",
		}, "\r\n"))

		for _, content := range contents {
			bytes := []byte(content)
			fmt.Fprintf(conn, "%x\r\n%s\r\n", len(bytes), content)
		}
		fmt.Fprintf(conn, "0\r\n\r\n")
	}
}

func main() {
	addr := "localhost:8888"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server is runniing at %v\n", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go processSession(conn)
	}
}
