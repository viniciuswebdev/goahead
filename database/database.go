package database

import (
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConf struct {
	User     string
	Password string
	Database string
}

type Database struct {
	User, Password, Database string
}

func Create(conf *DatabaseConf) *Database {
	db := new(Database)
	db.User = conf.User
	db.Password = conf.Password
	db.Database = conf.Database

	return db
}
