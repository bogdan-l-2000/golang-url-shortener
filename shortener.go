package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type UrlEntry struct {
	AliasUrl    string
	OriginalUrl string
}

var entry_MAP = map[string]UrlEntry{}

func main() {
	mux := http.NewServeMux()

	rh := http.HandlerFunc(addShortenedUrl)
	mux.HandleFunc("/addUrl", rh)

	sp := http.HandlerFunc(redirectUrl)
	mux.HandleFunc("/{shortPath}", sp)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func addShortenedUrl(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm error: %v", err)
			return
		}

		alias := r.FormValue("alias")
		original := r.FormValue("original")

		if alias == "" || original == "" {
			http.Error(w, "One or more form values are not present", http.StatusBadRequest)
		}

		if len(alias)+8 >= len(original) { // +8 refers to https:// prefix, this is a temporary validator
			http.Error(w, "The length of the alias must be shorter than the length of the original url", http.StatusBadRequest)
		}

		// urlRegex := `^(?:https?://)?(?:[^/.\s]+\.)*` + regexp.QuoteMeta(keyword) + `(?:/[^/\s]+)*/?$`
		// Source: https://stackoverflow.com/questions/67159622/golang-regexp-matchstring-handle-url-match

		urlRegex := `^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/|\/|\/\/)?[A-z0-9_-]*?[:]?[A-z0-9_-]*?[@]?[A-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`
		// Source: https://gist.github.com/brydavis/0c7da92bd508195744708eeb2b54ac96

		url_match, err := regexp.MatchString(urlRegex, original)
		if err != nil {
			http.Error(w, "Error when parsing the URL", http.StatusBadRequest)
		}
		if !url_match {
			http.Error(w, "The original value must be a URL format", http.StatusBadRequest)
		}

		fmt.Printf("%s\n", alias)
		fmt.Printf("%s\n", original)

		entry_MAP[alias] = UrlEntry{AliasUrl: alias, OriginalUrl: original}
	default:
		http.Error(w, "400 bad request", http.StatusBadRequest)
	}
	fmt.Printf("New path has been added.\n")
	w.WriteHeader(http.StatusCreated)
}

func redirectUrl(w http.ResponseWriter, r *http.Request) {
	if key, ok := entry_MAP[r.PathValue("shortPath")]; ok {
		http.Redirect(w, r, key.OriginalUrl, http.StatusTemporaryRedirect)
	} else {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	fmt.Printf("Redirected!\n")
}
