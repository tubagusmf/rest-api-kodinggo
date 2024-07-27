package db

import (
	"database/sql"
	"golang-rest-api-articles/internal/helper"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewMysql() *sql.DB {
	//koneksi database
	db, err := sql.Open("mysql", helper.GetConnectionString())
	if err != nil {
		log.Fatal(err)
	}

	//pengecekan koneksi database
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
