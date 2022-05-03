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

func IsFile(Id int) (bool, error) {
	sql, err := db.Prepare("SELECT `FileSize` FROM `tree` WHERE `Id` = ?")
	if err != nil {
		log.Panic(err)
	}
	defer sql.Close()
	rows, err := sql.Query(Id)
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()
	var temp int
	rows.Next()
	rows.Scan(&temp)
	return temp != -1, err
}

func GetFileName(Id int) (string, error) {
	sql, err := db.Prepare("SELECT `FileName` FROM `tree` WHERE `Id` = ? AND `FileSize` <> -1")
	if err != nil {
		log.Panic(err)
	}
	defer sql.Close()
	rows, err := sql.Query(Id)
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()
	var temp string
	rows.Next()
	err = rows.Scan(&temp)
	return "./data/" + temp, err
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

func CreateFile(temp Tree) error {
	sql, _ := db.Prepare("INSERT INTO `tree` (`Id`, `Name`, `FileName`, `FileSize`, `Parent`, `Uptime`) VALUES (NULL, ?, ?, ?, ?, ?);")
	defer sql.Close()
	_, err = sql.Exec(temp.Name, temp.FileName, temp.FileSize, temp.Parent, temp.Uptime)
	return err
}

func UpdateCount(Id int, Count int) error {
	sql, _ := db.Prepare("UPDATE `tree` SET `FileSize` = ? WHERE `tree`.`Id` = ?;")
	defer sql.Close()
	_, err = sql.Exec(Count, Id)
	return err
}

func MoveFile(From int, To int) error {
	sql, _ := db.Prepare("UPDATE `tree` SET `Parent` = ? WHERE `tree`.`Id` = ?;")
	defer sql.Close()
	_, err = sql.Exec(To, From)
	return err
}

func DropFile(Id int) error {
	var flag int
	var p int
	sql, _ := db.Prepare("SELECT `FileSize`, `Parent` FROM `tree` WHERE `Id` = ?")
	rows, _ := sql.Query(Id)
	rows.Next()
	err = rows.Scan(&flag, &p)
	if err != nil {
		return err
	}
	if flag == -1 {
		sql, _ := db.Prepare("UPDATE `tree` SET `Parent` = ? WHERE `tree`.`Parent` = ?")
		defer sql.Close()
		_, err = sql.Exec(p, Id)
		sql, _ = db.Prepare("DELETE FROM `tree` WHERE `tree`.`Id` = ?")
		_, err = sql.Exec(Id)
	} else {
		sql, _ := db.Prepare("DELETE FROM `tree` WHERE `tree`.`Id` = ?")
		defer sql.Close()
		_, err = sql.Exec(Id)
	}
	return err
}
