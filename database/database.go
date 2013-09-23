package database

import (
	"errors"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConf struct {
	Driver,	User, Password, Name string
}

type Database struct {
	Driver, User, Password, Name string
}

type TableConf struct {
	Name, Hash, Url string
}

var drivers = []string{"mysql", "sqlite3"}

func (database *Database) FindShortenerUrlByHash(hash string, tableConf *TableConf) (string, error) {
	db, err := sql.Open(database.Driver, database.User+":"+database.Password+"@/"+database.Name)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var url string
	err = db.QueryRow("SELECT " + tableConf.Url + " FROM " + tableConf.Name + " WHERE " + tableConf.Hash + " = ?", hash).Scan(&url)
	if err != nil {
		return url, err
	}
	return url, nil
}

func (database *Database) IsValid() error{
	var isOk = false

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


func Create(conf *DatabaseConf) *Database {
	db := new(Database)
	db.Driver = conf.Driver 
	db.User = conf.User
	db.Password = conf.Password
	db.Name = conf.Name

	return db
}

