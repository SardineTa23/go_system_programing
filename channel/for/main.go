package main

import (
	"fmt"
	"math"
)

func primeNumber() chan int {
	// チャネル作成
	result := make(chan int)
	go func() {
		// チャネルに値入れる
		result <- 2
		// 3〜10000まで2ずつ加算していくfor
		for i := 3; i < 10000; i += 2 {
			//平方根を求める
			l := int(math.Sqrt(float64(i)))
			found := false
			for j := 3; j < l; j += 2 {
				found = true
				break
			}
			if !found {
				result <- 1
			}
		}
		close(result)
	}()
	return result
}

func main() {
	pn := primeNumber()
	for n := range pn {
		fmt.Println(n)
	}
}

// $  go run main.go
// 2
// 1
// 1
// 1
// 1
// 1
// 1
// 1
