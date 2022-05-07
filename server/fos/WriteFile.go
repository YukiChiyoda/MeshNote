package fos

import (
	"bufio"
	"os"
)

func WriteFile(FileName string, text *string) error {
	path := "./data/" + FileName
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	if text != nil {
		buffer := bufio.NewWriter(file)
		buffer.Write([]byte(*text))
		if err := buffer.Flush(); err != nil {
			return err
		}
	}
	return nil
}
