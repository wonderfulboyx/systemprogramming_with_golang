# メモ
## io.Readerの補助関数
- goでのメモリ確保はmakeを使う
    ```go
    // 1024バイトのバッファをmakeでつくる
    buffer := make([]byte, 1024)
    ```
- バッファの管理を考えると結構面倒くさいので、便利な補助機能がいくつかある
- `ioutil.ReadAll(r io.Reader) ([]byte, error)` でだいたいいける。
- 決まったバイト数だけ確実に読み込みたい場合は `io.ReadFull(r Reader, buf []byte) (n int, err error)`
- readerからwriterにコピーしたいときは`io.Copy(dst Writer, src Reader) (written int64, err error)` 。srcをdstに流す。
- 決まったバッファを使いまわしたいときは `io.CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)`

## ioインターフェースのキャスト
- 関数からは`io.ReadCloser`を要求されているが、単体テストではCloseメソッドを持たない構造体を使いたいケース
    -  `ioutil.NopCloser(r io.Reader) io.ReadCloser` にreaderをわたしてやると何もしないCloseをくっつけて返してくれる
- io.Readerとio.Writerをつなげるには `bufio.NewReadWriter(r *Reader, w *Writer) *ReadWriter`

## その他
- file.Openとfile.Createは内部的に同じ関数を呼んでいて、同じシステムコールが呼ばれている
```go
// Open opens the named file for reading. If successful, methods on
// the returned file can be used for reading; the associated file
// descriptor has mode O_RDONLY.
// If there is an error, it will be of type *PathError.
func Open(name string) (*File, error) {
	return OpenFile(name, O_RDONLY, 0)
}

// Create creates or truncates the named file. If the file already exists,
// it is truncated. If the file does not exist, it is created with mode 0666
// (before umask). If successful, methods on the returned File can
// be used for I/O; the associated file descriptor has mode O_RDWR.
// If there is an error, it will be of type *PathError.
func Create(name string) (*File, error) {
	return OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
}
```
