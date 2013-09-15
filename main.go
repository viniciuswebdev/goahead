package main

import (
	"./database"
	"code.google.com/p/gcfg"
	"net/http"
    "flag"
)

type Config struct {
	Mysql struct {
		User     string
		Password string
		Database string
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
    configFilePath := flag.String("config", "goahead.cfg", "Configuration file path")
    flag.Parse()

    var cfg Config
    gcfg.ReadFileInto(&cfg, *configFilePath)

	db := database.Database{cfg.Mysql.User, cfg.Mysql.Password, cfg.Mysql.Database}
	http.Redirect(w, r, db.FindShortenerUrlByHash(r.URL.Path[1:]), 301)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
