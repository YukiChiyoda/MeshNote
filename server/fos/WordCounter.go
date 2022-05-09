package fos

import "unicode/utf8"

func WordCount(text *string) int {
	result := utf8.RuneCountInString((*text))
	return result
}
