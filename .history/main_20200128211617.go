package main

import (
	"net/http"
	"net/ro"
)

func main() {
	buildHandler := http.FileServer(http.Dir("./client/build"))
	router.PathPrefix("/").Handler(buildHandler)
}
