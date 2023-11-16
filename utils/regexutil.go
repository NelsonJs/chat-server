package utils

import "regexp"

func IsImage(mimeType string) bool {
	isMatch, _ := regexp.MatchString("^image/[A-Za-z]{3,4,5}$", mimeType)
	return isMatch
}
