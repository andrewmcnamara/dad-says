package main

import (
	"net/http"
	"rour"
)

func main() {
	buildHandler := http.FileServer(http.Dir("./client/build"))
	router.PathPrefix("/").Handler(buildHandler)
}
