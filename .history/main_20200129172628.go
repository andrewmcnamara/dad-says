package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	buildHandler := http.FileServer(http.Dir("./client/build"))
	mux.Handle("/stuff", http.StripPrefix("/static/", fs))
	// mux.HandleFunc("/", buildHandler)
	// rh := http.RedirectHandler("http://example.org", 307)
	// mux.Handle("/foo", rh)

	log.Println("Listening...")
	http.ListenAndServe(":8000", mux)
	//
	// router.PathPrefix("/").Handler(buildHandler)
}
