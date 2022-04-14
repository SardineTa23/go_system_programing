package main

import (
	"bufio"
	"os"
)

func main() {
	// 出力結果を一時的にためておいて、ある程度の分量ごとにまとめて書き出す
	buffer := bufio.NewWriter(os.Stdout)

	buffer.WriteString("hogehoge ")
	// Flush() メソッドを呼ぶと、後続の io.Writer に書き出す
	buffer.Flush()

	buffer.WriteString("bufio.Writer\n")
	buffer.Flush()
}
