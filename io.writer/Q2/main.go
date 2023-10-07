package main

import (
	"encoding/csv"
	"os"
)

func main() {
	f, err := os.Create("MultiWriter.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	cw := csv.NewWriter(f)
	cw.Write([]string{"hoge", "fuga", "piyo"})
	cw.Flush()
}
