package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

// 改行/単語で区切る

var source = `1行め
2行め
3行め`

func main() {
	b := bytes.NewReader([]byte(source))
	br := bufio.NewReader(b)
	for {
		// 改行で区切る
		line, err := br.ReadString('\n')
		fmt.Printf("%#v\n", line)
		if err == io.EOF {
			break
		}
	}
}

// $ go run main.go
// "1行め\n"
// "2行め\n"
// "3行め"
