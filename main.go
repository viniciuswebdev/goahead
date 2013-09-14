package main

import (
    "./database"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    r1 := database.Database{"root", "abc123", "zocprint"}
    http.Redirect(w, r, r1.FindShortenerUrlByHash("zocprint"), 301)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}