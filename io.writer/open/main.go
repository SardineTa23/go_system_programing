package main

import "os"

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	file.Write([]byte("hogehoge os.File Create!\n"))
	file.Close()
}
