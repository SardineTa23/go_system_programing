package main

import (
	"io"
	"os"
)

func main() {
	f, err := os.Create("MultiWriter.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := io.MultiWriter(f, os.Stdout)
	w.Write([]byte("標準出力とファイル､同時に書き込めます\n"))
}
