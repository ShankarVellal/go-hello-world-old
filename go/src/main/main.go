package main

import (
	"fmt"
	"net/http"
)

func indexHandler (w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-Forwarded-Proto") == "https" {
		fmt.Fprintf(w, "Hello, World! Secure!")
	} else {	
		http.Redirect(w, r, "https://shankarvellal.com:443"+r.RequestURI, http.StatusMovedPermanently)
	}
}

func elbHealthCheck (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Health OK!")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/elb_health_check", elbHealthCheck)
	http.ListenAndServe(":8080", nil)
}