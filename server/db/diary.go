package db

import (
	_ "github.com/go-sql-driver/mysql"
)

func QueryDiary(data Diary) (int, error) {
	res, err := db.Preparex("SELECT `count` FROM `diary` WHERE `user` = ? AND `year` = ? AND `month` = ? AND `day` = ? LIMIT 1")
	if err != nil {
		return SPECIAL_ERROR_ID, err
	}
	temp := DIARY_EMPTY
	res.Get(&temp, data.User, data.Year, data.Month, data.Day) // No row will return 0 conut
	return temp, nil
}

func UpdateDiary(data Diary) error {
	temp, err := QueryDiary(data)
	if err != nil {
		return err
	}
	data.Count += temp
	if temp == DIARY_EMPTY {
		res, err := db.Preparex("INSERT INTO `diary` (`id`, `user`, `year`, `month`, `day`, `count`) VALUES (NULL, ?, ?, ?, ?, ?)")
		if err != nil {
			return err
		}
		_, err = res.Exec(data.User, data.Year, data.Month, data.Day, data.Count)
		if err != nil {
			return err
		}
	} else {
		res, err := db.Preparex("UPDATE `diary` SET `count` = ? WHERE `user` = ? AND `year` = ? AND `month` = ? AND `day` = ?")
		if err != nil {
			return err
		}
		_, err = res.Exec(data.Count, data.User, data.Year, data.Month, data.Day)
		if err != nil {
			return err
		}
	}
	return nil
}
