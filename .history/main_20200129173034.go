package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./client/build"))
	mux.Handle("/", http.StripPrefix("/client/", fs))

	staticAssetsFs := http.StripPrefix("/static/", http.FileServer(http.Dir("./build/static")))
	mux.Handle("/static/").Handler(staticAssetsFs)
	// mux.HandleFunc("/", buildHandler)
	// rh := http.RedirectHandler("http://example.org", 307)
	// mux.Handle("/foo", rh)

	log.Println("Listening...")
	http.ListenAndServe(":8000", mux)
	//
	// router.PathPrefix("/").Handler(buildHandler)
}
