package main

import (
	"fmt"
	"log"
	"net/http"
)

type UrlEntry struct {
	AliasUrl    string
	OriginalUrl string
}

var entry_MAP = map[string]UrlEntry{}

func addShortenedUrl(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm error: %v", err)
			return
		}
		alias := r.FormValue("alias")
		original := r.FormValue("original")
		fmt.Printf("%s", alias)
		fmt.Printf("%s", original)

		entry_MAP[alias] = UrlEntry{AliasUrl: alias, OriginalUrl: original}
	default:
		http.Error(w, "400 bad request", http.StatusBadRequest)
	}
	fmt.Printf("New path has been added.")
}

func redirectUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Redirected!")
}

func main() {
	mux := http.NewServeMux()

	rh := http.HandlerFunc(addShortenedUrl)
	mux.HandleFunc("/addUrl", rh)

	sp := http.HandlerFunc(redirectUrl)
	mux.HandleFunc("/{shortPath}", sp)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
