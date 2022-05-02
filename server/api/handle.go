package api

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Tree struct {
	Id       int    `json:"Id"`
	Name     string `json:"Name"`
	FileName string `json:"FileName"`
	FileSize int    `json:"FileSize"` // -1 => dir
	Parent   int    `json:"Parent"`   // -1 => /
	Uptime   string `json:"Uptime"`
}

var db *sql.DB
var err error

func init() {
	if db, err = sql.Open("mysql", "mesh:mesh@tcp(localhost:3306)/meshnote"); err != nil {
		log.Panic(err)
	}
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	//defer db.Close()
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	log.Println("MySQL connected successfully!")
}

func QueryFile(Parent int, Data *[]Tree) {
	sql, err := db.Prepare("SELECT * FROM `tree` WHERE `Parent` = ?")
	if err != nil {
		log.Panic(err)
	}
	defer sql.Close()
	rows, err := sql.Query(Parent)
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()
	for i := 0; rows.Next(); i++ {
		var temp Tree
		if err := rows.Scan(&temp.Id, &temp.Name, &temp.FileName, &temp.FileSize, &temp.Parent, &temp.Uptime); err != nil {
			log.Panic(err)
		}
		*Data = append(*Data, temp)
	}
	log.Printf("Query all files in Parent[%d] successfully!\n", Parent)
}

func QueryAllFile(Data *[]Tree) {
	sql, err := db.Prepare("SELECT * FROM `tree`")
	if err != nil {
		log.Panic(err)
	}
	defer sql.Close()
	rows, err := sql.Query()
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()
	for i := 0; rows.Next(); i++ {
		var temp Tree
		if err := rows.Scan(&temp.Id, &temp.Name, &temp.FileName, &temp.FileSize, &temp.Parent, &temp.Uptime); err != nil {
			log.Panic(err)
		}
		*Data = append(*Data, temp)
	}
	log.Println("Query all files in database successfully!")
}

func CreateFile(temp Tree) {

}

func UpdateCount(FileName string, Count int) {

}
