// Operate MySQL
package db

import (
	"MeshNote/server/fos"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Connect("mysql", "mesh:mesh@tcp(localhost:3306)/meshnote")
	handleError(err)
	db.SetMaxOpenConns(10)
	db.SetMaxOpenConns(10)
}

func handleError(err error) {
	if err != nil {
		log.Println(err.Error())
		fos.UpdateErrorLog()
	}
}

func QueryFile(parent int) ([]Tree, error) {
	res, err := db.Preparex("SELECT * FROM `tree` WHERE `parent` = ?")
	if err != nil {
		handleError(err)
		return nil, err
	}
	var temp []Tree
	err = res.Select(&temp, parent)
	if err != nil {
		handleError(err)
		return nil, err
	}
	return temp, nil
}

func CreateFolder(t Tree) error {
	//res, err := db.Preparex("INSERT INTO `tree` VALUES (?, ?, ?, ?, ?, ?, ?)")
	return nil
}
