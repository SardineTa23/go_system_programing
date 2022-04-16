package main

import (
	"fmt"
	"strings"
)

// データ型を指定して解析

var source = "123 1.234 1.0e4 test"

func main() {
	reader := strings.NewReader(source)
	var i int
	var f, g float64
	var s string
	fmt.Fscan(reader, &i, &f, &g, &s)
	fmt.Printf("i=%v f=%v g=%v  s=%v", i, f, g, s)
}

// $ go run main.go
// i=123 f=1.234 g=10000  s="test"
