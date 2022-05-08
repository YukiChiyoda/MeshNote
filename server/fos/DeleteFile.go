package fos

import "os"

func DeleteFile(fileName string) error {
	path := "./data/" + fileName
	return os.Remove(path)
}
