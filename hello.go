package main

import (
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	basepath := "http://www.google.com.br/"
    http.Redirect(w, r, basepath + r.URL.Path[1:], 301)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}