package db

import (
	_ "github.com/go-sql-driver/mysql"
)

func QueryDiary(user int, year int, month int, day int) (int, error) {
	res, err := db.Preparex("SELECT `count` FROM `diary` WHERE `user` = ? AND `year` = ? AND `month` = ? AND `day` = ? LIMIT 1")
	if err != nil {
		return SPECIAL_ERROR_ID, err
	}
	temp := DIARY_EMPTY
	res.Get(&temp, user, year, month, day) // No row will return 0 conut
	return temp, nil
}

func UpdateDiary(user int, year int, month int, day int, count int) error {
	temp, err := QueryDiary(user, year, month, day)
	if err != nil {
		return err
	}
	count += temp
	if temp == DIARY_EMPTY {
		res, err := db.Preparex("INSERT INTO `diary` (`id`, `user`, `year`, `month`, `day`, `count`) VALUES (NULL, ?, ?, ?, ?, ?)")
		if err != nil {
			return err
		}
		_, err = res.Exec(user, year, month, day, count)
		if err != nil {
			return err
		}
	} else {
		res, err := db.Preparex("UPDATE `diary` SET `count` = ? WHERE `user` = ? AND `year` = ? AND `month` = ? AND `day` = ?")
		if err != nil {
			return err
		}
		_, err = res.Exec(count, user, year, month, day)
		if err != nil {
			return err
		}
	}
	return nil
}
