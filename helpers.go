package nbc

import (
	"regexp"
)

var reg = regexp.MustCompile("\\s+")

func SplitText(text string) []string {
	return reg.Split(text, -1)
}
