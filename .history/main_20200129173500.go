package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./client/build"))
	mux.Handle("/", fs)

	staticAssetsFs := http.StripPrefix("/static/", http.FileServer(http.Dir("./client/build/static")))
	mux.Handle("/static/", staticAssetsFs)

	log.Println("Listening...")
	http.ListenAndServe(":8000", mux)
}
