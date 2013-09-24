package database

import (
	"errors"
	"database/sql"
)

type TableConf struct {
	Name, Hash, Url string
}

func (database *Database) FindShortenerUrlByHash(hash string, tableConf *TableConf) (string, error) {
	db, err := sql.Open(database.Driver, database.dataSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var url string
	err = db.QueryRow("SELECT " + tableConf.Url + " FROM " + tableConf.Name + " WHERE " + tableConf.Hash + " = $1 ", hash).Scan(&url)
	if err != nil {
		return url, err
	}
	return url, nil
}

func (database *Database) IsValid() error{
	var isOk = false
	db, err := sql.Open(database.Driver, database.dataSource)
	if err != nil {
		return err 
	}
	defer db.Close()
	err = db.Ping() 
	if err != nil {
		return err 
	}

	for _, driver:= range drivers{
		if database.Driver == driver {
			isOk = true 
			break
		}
	}

	if !isOk {
		return errors.New("Driver "+database.Driver+" is not supported.")
	}

	return nil 
}