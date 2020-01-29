package main

import (
	"net/http"
)

func main() {
	buildHandler := http.FileServer(http.Dir("./))
	router.PathPrefix("/").Handler(buildHandler)
}
