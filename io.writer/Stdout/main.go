package main

import (
	"os"
)

func main() {
	os.Stdout.Write([]byte("hogehoge os.Stdout\n"))
}
