package main

import (
	"fmt"
	"net/http"
)

func indexHandler (w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-Forwarded-Proto") == "https" {
		fmt.Fprintf(w, "Hello, World!")
	}
	else {	
		http.Redirect(w, r, "https://shankarvellal.com"+r.RequestURI, http.StatusMovedPermanently)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}