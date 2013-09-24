package database

import (
	"database/sql"
	"github.com/pmylund/go-cache"
	"time"
)

type TableConf struct {
	Name, Hash, Url string
}

func (database *Database) FindShortenerUrlByHash(hash string, tableConf *TableConf) (string, error) {
	db, err := sql.Open("mysql", database.User+":"+database.Password+"@/"+database.Database)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	c := cache.New(5*time.Minute, 30*time.Second)
	cachedUrl, found := c.Get(tableConf.Hash)
	if found {
        return cachedUrl.(string), nil
	}

	var url string
	err = db.QueryRow("SELECT " + tableConf.Url + " FROM " + tableConf.Name + " WHERE " + tableConf.Hash + " = ?", hash).Scan(&url)
	if err != nil {
		return url, err
	}

	c.Set(tableConf.Hash, url, 0)

	return url, nil
}