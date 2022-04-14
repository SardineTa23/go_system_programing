package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	source := map[string]string{
		"Hello": "World",
	}

	file, _ := os.Create("hoge.json.gz")
	gzip := gzip.NewWriter(file)

	multiWriter := io.MultiWriter(gzip, os.Stdout)

	encoder := json.NewEncoder(multiWriter)
	encoder.SetIndent("", "   ")
	err := encoder.Encode(source)
	if err != nil {
		log.Fatal(err)
		return
	}

	gzip.Flush()
	gzip.Close()
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
