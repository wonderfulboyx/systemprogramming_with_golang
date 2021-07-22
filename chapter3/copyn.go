package main

import (
	"io"
)

func CopyN(dest io.Writer, src io.Reader, length int) {
	r := io.LimitReader(src, int64(length))
	io.Copy(dest, r)
}
