package main

import (
	"net/http"
	
)

func main() {
	buildHandler := http.FileServer(http.Dir("./client/build"))
	router.PathPrefix("/").Handler(buildHandler)
}
