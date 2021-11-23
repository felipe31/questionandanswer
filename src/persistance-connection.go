package main

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func MysqlConnect(host string) error {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "123456123456",
		Net:                  "tcp",
		DBName:               "questionandanswer",
		AllowNativePasswords: true,
		Addr:                 host + ":3306",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return pingErr
	}
	fmt.Println("Connected!")
	return nil
}
