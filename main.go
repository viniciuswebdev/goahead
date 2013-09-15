package main

import (
	"./database"
	"code.google.com/p/gcfg"
	"net/http"
    "flag"
    "log"
)



type Config struct {
    General struct {
        Port string 
    }
    Mysql database.DatabaseConf 
}

var _cfg Config

func handler(w http.ResponseWriter, r *http.Request) {
	db := database.Create(&(_cfg.Mysql))
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
	log.Printf("Starting Goahead on localhost:%s ...\n", _cfg.General.Port)
    err = http.ListenAndServe(":"+_cfg.General.Port, nil)
    if err != nil {
        panic(err.Error())
    }
}
