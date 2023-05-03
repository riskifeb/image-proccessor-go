package main

import (
	"log"
	"net/http"

	"github.com/riskifeb/compresGambar/handler"
)

func main() {
	http.HandleFunc("/upload", handler.HandleUpload)
	log.Fatal(http.ListenAndServe(":5005", nil))
}
