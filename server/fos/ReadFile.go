package fos

import (
	"io"
	"os"
)

func ReadFile(fileName string) (string, error) {
	path := "./data/" + fileName
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	var temp string
	buffer := make([]byte, 100)
	for {
		reader, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				return "", err
			}
			break
		}
		temp += string(buffer[:reader])
	}
	return temp, nil
}
