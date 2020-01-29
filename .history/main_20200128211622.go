package main

import (
	"net/http"
	"rou"
)

func main() {
	buildHandler := http.FileServer(http.Dir("./client/build"))
	router.PathPrefix("/").Handler(buildHandler)
}
