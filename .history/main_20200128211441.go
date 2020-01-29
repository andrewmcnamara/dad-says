
package main 

main() {
	buildHandler := http.FileServer(http.Dir("<path to build>"))
	router.PathPrefix("/").Handler(buildHandler)
}
