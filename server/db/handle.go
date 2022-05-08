// Operate MySQL
package db

import (
	"errors"
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
	if id == PARENT_ROOT {
		return TYPE_DIR, nil
	}
	if id < PARENT_ROOT {
		return SPECIAL_INVALID_ID, errors.New("db: invalid id")
	}
	res, err := db.Preparex("SELECT `type` FROM `tree` WHERE `id` = ?")
	if err != nil {
		return SPECIAL_ERROR_ID, err
	}
	var temp int
	err = res.Get(&temp, id)
	if err != nil {
		return SPECIAL_ERROR_ID, err
	}
	return temp, nil
}

func GetElementFileName(id int) (string, error) {
	res, err := db.Preparex("SELECT `filename` FROM `tree` WHERE `id` = ?")
	if err != nil {
		return SPECIAL_NULL_NAME, err
	}
	var temp string
	err = res.Get(&temp, id)
	if err != nil {
		return SPECIAL_NULL_NAME, err
	}
	return temp, nil
}

func GetElementParent(id int) (int, error) {
	if id == PARENT_ROOT {
		return TYPE_DIR, nil
	}
	res, err := db.Preparex("SELECT `parent` FROM `tree` WHERE `id` = ?")
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
	if data.Type == TYPE_FILE || data.Type == TYPE_RECYLED_FILE {
		_, err = res.Exec(data.Name, data.Type, data.FileName, data.FileSize, data.Parent, data.Uptime)
	} else {
		_, err = res.Exec(data.Name, data.Type, FILENAME_DIR, data.FileSize, data.Parent, data.Uptime)
	}
	if err != nil {
		return err
	}
	return nil
}

func MoveElement(id int, target int) error {
	res, err := db.Preparex("UPDATE `tree` SET `parent` = ? WHERE `id` = ?")
	if err != nil {
		return err
	}
	_, err = res.Exec(target, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteElement(id int) error {
	fileType, err := GetElementType(id)
	if err != nil {
		return err
	}
	switch fileType {
	case TYPE_FILE:
		res, err := db.Preparex("UPDATE `tree` SET `type` = ? WHERE `id` = ?")
		if err != nil {
			return err
		}
		_, err = res.Exec(TYPE_RECYLED_FILE, id)
		if err != nil {
			return err
		}
		return nil
	case TYPE_RECYLED_FILE:
		return errors.New("delete: this file has already been recycled")
	case TYPE_DIR:
		sql, err := db.Beginx()
		if err != nil {
			return err
		}
		res, err := sql.Preparex("UPDATE `tree` SET `parent` = ? WHERE `parent` = ?")
		if err != nil {
			sql.Rollback()
			return err
		}
		newParent, err := GetElementParent(id)
		if err != nil {
			sql.Rollback()
			return err
		}
		_, err = res.Exec(newParent, id)
		if err != nil {
			sql.Rollback()
			return err
		}
		res, err = sql.Preparex("DELETE FROM `tree` WHERE `id` = ?")
		if err != nil {
			sql.Rollback()
			return err
		}
		_, err = res.Exec(id)
		if err != nil {
			sql.Rollback()
			return err
		}
		err = sql.Commit()
		if err != nil {
			sql.Rollback()
			return err
		}
		return nil
	}
	return errors.New("delete: type undefined error")
}
