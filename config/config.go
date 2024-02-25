package config

import "database/sql"

var db *sql.DB

func Init() error{
	var err error
	db, err = InitializeMySQL();

	if err != nil {
		panic(err)
	}

	return nil

}

func GetMysql() *sql.DB {
	return db
}