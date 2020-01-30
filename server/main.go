package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Joke struct {
	ID   string `json:"id"`
	Text string `json:"joke"`
}

func fetchJoke() Joke {

	url := "https://icanhazdadjoke.com/"
	jokeClient := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "My Library (https://github.com/andrewmcnamara/dad-jokes)")
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, getErr := jokeClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	joke := Joke{}
	jsonErr := json.Unmarshal(body, &joke)
	if jsonErr != nil {
		fmt.Println(string(body))
		log.Fatal(jsonErr)
	}

	return joke
}

// Logger records incoming requests and uses the log package to print http method,
// route url and duration to screen/logs.
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Save start time to calculate duration
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t\t%s\t\t%s\t\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

func main() {

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./client/build"))
	mux.Handle("/", fs)

	staticAssetsFs := http.StripPrefix("/static/", http.FileServer(http.Dir("./client/build/static")))
	mux.Handle("/static/", staticAssetsFs)

	jokeHandler := func(w http.ResponseWriter, req *http.Request) {
		joke := fetchJoke()
		fmt.Println()
		io.WriteString(w, joke.Text+"\n")
	}

	mux.HandleFunc("/joke", jokeHandler)
	log.Println("Listening...")
	http.ListenAndServe(":8000", Logger(mux, "WEB"))
}
