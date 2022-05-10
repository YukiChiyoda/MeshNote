package fos

import (
	"log"
	"os"
	"time"
)

func UpdateErrorLog(errorMsg error) {
	path := "./error.log"
	temp := time.Now().Format("2006-01-02 15:04:05 ") + errorMsg.Error() + "\n"
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Panicln(err.Error())
		return
	}
	defer file.Close()
	_, err = file.WriteString(temp)
	if err != nil {
		log.Panicln(err.Error())
		return
	}
}
