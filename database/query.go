package database

import (
	"database/sql"
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

	var url string
	err = db.QueryRow("SELECT " + tableConf.Url + " FROM " + tableConf.Name + " WHERE " + tableConf.Hash + " = ?", hash).Scan(&url)
	if err != nil {
		// no rows matched! Returns 404
		return url, err
	}
	return url, nil
}