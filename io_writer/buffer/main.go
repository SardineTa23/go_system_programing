package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("hogehoge byte.Buffer!\n"))
	fmt.Println(buffer.String())
}
