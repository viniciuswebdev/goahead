package main

import (
    "./database"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    db := database.Database{"root", "abc123", "zocprint"}
    http.Redirect(w, r, db.FindShortenerUrlByHash(r.URL.Path[1:]), 301)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}