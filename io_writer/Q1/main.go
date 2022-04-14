package main

import (
	"fmt"
	"os"
)

//fmt. Fprintf(writer, フォーマット文字列, 値...)でio.Writerに数値や文字列を 出力できます。fmt.Fprintf() では、他の言語と同じように、"%d" が数値の表示 に、"%s" が文字列の表示に、"%f" が浮動小数点数の表示に使えます。
//  これらを使って、数字や小数、文字列をファイルに書き出してみましょう。

func main() {
	fmt.Fprintf(os.Stdout, "%d %s %f", 3, "fefe", 2.4)
}
