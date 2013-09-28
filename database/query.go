package database

import (
	"errors"
	"database/sql"
	"github.com/pmylund/go-cache"
	"time"
	"fmt"
)

type TableConf struct {
	Name, Hash, Url string
}

func (database *Database) FindShortenedUrlByHash(hash string, tableConf *TableConf) (string, error) {
	db, err := sql.Open(database.Driver, database.dataSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	c := cache.New(15*time.Minute, 30*time.Second)
	cachedUrl, found := c.Get(tableConf.Hash)
	if found {
        return cachedUrl.(string), nil
	}

	database_var := "?"
	if database.IsPostgres() {
		database_var = "$"
	}
	var url string
	err = db.QueryRow(fmt.Sprintf("SELECT %s FROM %s WHERE %s = %s", tableConf.Url, tableConf.Name, tableConf.Hash, database_var), hash).Scan(&url)
	if err != nil {
		return url, err
	}
	if len(url) != 0 {
		c.Set(tableConf.Hash, url, 0)
	}
	return url, nil
}
