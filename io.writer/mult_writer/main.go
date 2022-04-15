package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}

	wirter := io.MultiWriter(file, os.Stdout)
	io.WriteString(wirter, "hogehoge io,MultiWriter\n")
}
