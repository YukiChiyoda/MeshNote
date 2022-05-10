package db

import (
	"log"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Connect("mysql", "mesh:mesh@tcp(localhost:3306)/meshnote")
	if err != nil {
		log.Panicln(err)
	}
	// defer db.Close() [Warning: This will cause `sql: database is closed`!]
	db.SetMaxOpenConns(10)
	db.SetMaxOpenConns(10)
}
