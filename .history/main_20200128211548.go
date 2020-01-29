package main

import (
	"net/http"
)

func main() {
	buildHandler := http.FileServer(http.Dir("./client))
	router.PathPrefix("/").Handler(buildHandler)
}
