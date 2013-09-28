package database

import (
	"log"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
)
var drivers = []string{"mysql", "sqlite3", "postgres"}

type DatabaseConf struct {
	Driver,	User, Password, Name, Path string
}

type Database struct {
	Driver,	User, Password, Name, Path, dataSource string
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

func (database *Database) IsPostgres() bool {
	return database.Driver == "postgres"
}

func (database *Database) IsSqlite3() bool {
	return database.Driver == "sqlite3"
}

func (database *Database) IsMysql() bool {
	return database.Driver == "mysql"
}