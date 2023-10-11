package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	source := map[string]string{
		"Hello": "World",
	}

	f, _ := os.Create("hoge.json.gz")
	defer f.Close()
	gf := gzip.NewWriter(f)
	defer gf.Close()

	multiWriter := io.MultiWriter(gf, os.Stdout)

	encoder := json.NewEncoder(multiWriter)
	encoder.SetIndent("", "   ")
	err := encoder.Encode(source)
	if err != nil {
		log.Fatal(err)
		return
	}

	gf.Flush()
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
