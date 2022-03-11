package config

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func NewSqlite3Database(config Config) *sql.DB {
	db, err := sql.Open("sqlite3", config.Get("FILE_PATH"))
	if err != nil {
		logrus.Panic(err)
	}

	if db == nil {
		logrus.Panic("db is nill")
	}

	return db

}
