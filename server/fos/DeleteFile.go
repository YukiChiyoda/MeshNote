package fos

import "os"

func DeleteFile(fileName string) error {
	path := "./data/" + fileName
	name := "./data/" + "." + fileName
	// return os.Remove(path)
	return os.Rename(path, name)
}
