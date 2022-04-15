package main

import (
	"io"
	"os"
)

func main() {
	// ファイルからの入力には、os.File 構造体を使う。
	file, err := os.Open("file.go")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	io.Copy(os.Stdout, file)
}
