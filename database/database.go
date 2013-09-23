package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConf struct {
	Driver string 
	User     string
	Password string
	Name string
}

type Database struct {
	Driver, User, Password, Name string
}

type TableConf struct {
	Name, Hash, Url string
}

func (database *Database) FindShortenerUrlByHash(hash string, tableConf *TableConf) (string, error) {
	db, err := sql.Open(database.Driver, database.User+":"+database.Password+"@/"+database.Name)
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

func Create(conf *DatabaseConf) *Database {
	db := new(Database)
	db.Driver = conf.Driver 
	db.User = conf.User
	db.Password = conf.Password
	db.Name = conf.Name

	return db
}
