package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

func NewMysqlDatabase() *sql.DB {

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASS"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_NAME"),
	)

	db, err := sql.Open("mysql", url)
	if err != nil {
		logrus.Panic(err)
	}

	err = db.Ping()
	if err != nil {
		logrus.Panic(err)
	}

	return db
}
