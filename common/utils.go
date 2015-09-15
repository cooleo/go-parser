package common

import (
   // "fmt"
	//"net/http"
	//"reflect"
	"strings"
	"unicode"
)

func GenerateTextSlug(Text string) string {
    //return utils.GenerateSlug(Text)
    return strings.Map(func(r rune) rune {
		switch {
		case r == ' ', r == '-':
			return '-'
		case r == '_', unicode.IsLetter(r), unicode.IsDigit(r):
			return r
		default:
			return -1
		}
		return -1
	}, strings.ToLower(strings.TrimSpace(Text)))
}