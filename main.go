package main

import (
	"code.google.com/p/gcfg"
	"flag"
	"github.com/viniciuswebdev/goahead/database"
	"github.com/viniciuswebdev/goahead/server"
)

type Config struct {
	General server.Server
	Database database.DatabaseConf
	Table database.TableConf
	Cache database.CacheConf
}

var _cfg Config
func main() {
	configFilePath := flag.String("config", "./etc/goahead.ini", "Configuration file path")
	flag.Parse()

	err := gcfg.ReadFileInto(&_cfg, *configFilePath)
	if err != nil {
		panic(err.Error())
	}
	db := database.Create(&(_cfg.Database))
	err = db.IsValid() 
	if err != nil {
		panic(err.Error())
	}
	_cfg.General.TurnOn(db, &(_cfg.Table), &(_cfg.Cache))
}
