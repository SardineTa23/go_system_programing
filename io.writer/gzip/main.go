package main

import (
	"compress/gzip"
	"io"
	"os"
)

func main() {
	f, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := gzip.NewWriter(f)
	defer w.Close()
	w.Header.Name = "test.txt"
	io.WriteString(w, "hogehoge, gzip.Writer\n")
}
