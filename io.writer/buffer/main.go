package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	// bytes.Bufferに書き込む
	buffer.Write([]byte("bytes.Buffer! \n"))
	//　書き込まれた内容を文字列として取得
	fmt.Println(buffer.String())
}
