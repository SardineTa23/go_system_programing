package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// 改行/単語で区切る

var source = `1行め
2行め
3行め`

func main() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
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
