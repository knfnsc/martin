package main

import (
	"log"
	"net/http"
)

const Port = ":8080"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", ShowHomePage)
	mux.HandleFunc("/write/{$}", ShowEditor)

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	log.Printf("✨ Ao vivo em \"%s\".\n", Port)
	log.Fatal(http.ListenAndServe(Port, mux))
}
