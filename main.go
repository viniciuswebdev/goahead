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

var _cfg Config

func handler(w http.ResponseWriter, r *http.Request) {
	db := database.Database{_cfg.Mysql.User, _cfg.Mysql.Password, _cfg.Mysql.Database}
	http.Redirect(w, r, db.FindShortenerUrlByHash(r.URL.Path[1:]), 301)
}

func main() {
    configFilePath := flag.String("config", "./etc/goahead.ini", "Configuration file path")
    flag.Parse()
    
    err := gcfg.ReadFileInto(&_cfg, *configFilePath)
    if err != nil {
        panic(err.Error())
    }

	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
