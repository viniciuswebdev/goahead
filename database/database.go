package database

import (
	"errors"
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
)

type DatabaseConf struct {
	Driver,	User, Password, Name, Path string
}

type Database struct {
	Driver, User, Password, Name, Path, dataSource string
}

type TableConf struct {
	Name, Hash, Url string
}

var drivers = []string{"mysql", "sqlite3", "postgres"}

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

func Create(conf *DatabaseConf) *Database {
	db := new(Database)
	db.Driver = conf.Driver 
	db.User = conf.User
	db.Password = conf.Password
	db.Name = conf.Name
	db.Path = conf.Path

	var dataSource string 
	if db.Driver == "sqlite3"{
		dataSource = db.Path
	} else if db.Driver == "postgres" {
		dataSource = "user="+db.User+" password="+db.Password+" dbname="+db.Name
	} else {
		dataSource = db.User+":"+db.Password+"@/"+db.Name
	}
	db.dataSource = dataSource

	log.Printf("Preparing connection with '%s' driver to %s\n", db.Driver, db.dataSource)
	return db
}

