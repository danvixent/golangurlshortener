package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", mapping)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func mapping(w http.ResponseWriter, r *http.Request) {
	mapper := make(map[string]string)
	mapper["/"] = "https://google.com"
	mapper["/base"] = "https://github.com"
	mapper["/go"] = "https://golang.org"
	mapper["/stack"] = "https://stackoverflow.com"
	mapper["/code"] = "https://codein.withgoogle.com"

	uri := r.URL.Path

	for key, val := range mapper {
		if uri == key {
			http.Redirect(w, r, val, http.StatusSeeOther)
			return
		}
	}
}
