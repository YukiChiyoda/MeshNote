// Operate MySQL
package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
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

func GetElementType(id int) (int, error) {
	if id == Parent_Root {
		return Type_Dir, nil
	}
	res, err := db.Preparex("SELECT `type` FROM `tree` WHERE `id` = ?")
	if err != nil {
		return -999, err
	}
	var temp int
	err = res.Get(&temp, id)
	if err != nil {
		return -999, err
	}
	return temp, nil
}

func GetElementFileName(id int) (string, error) {
	res, err := db.Preparex("SELECT `filename` FROM `tree` WHERE `id` = ?")
	if err != nil {
		return "null", err
	}
	var temp string
	err = res.Get(&temp, id)
	if err != nil {
		return "null", err
	}
	return temp, nil
}

func QueryElement(parent int) ([]Tree, error) {
	res, err := db.Preparex("SELECT * FROM `tree` WHERE `parent` = ?")
	if err != nil {
		return nil, err
	}
	var temp []Tree
	err = res.Select(&temp, parent)
	if err != nil {
		return nil, err
	}
	return temp, nil
}

func CreateElement(data Tree) error {
	res, err := db.Preparex("INSERT INTO `tree` (`id`, `name`, `type`, `filename`, `filesize`, `parent`, `uptime`) VALUES (NULL, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	if data.Type == Type_File || data.Type == Type_Recyled_File {
		_, err = res.Exec(data.Name, data.Type, data.FileName, data.FileSize, data.Parent, data.Uptime)
	} else {
		_, err = res.Exec(data.Name, data.Type, FileName_Dir, data.FileSize, data.Parent, data.Uptime)
	}
	if err != nil {
		return err
	}
	return nil
}
