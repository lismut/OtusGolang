package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(pack string) (string, error) {
	r := []rune(pack)
	res := ""
	for i, c := range r {
		if i == 0 && unicode.IsDigit(c) {
			return "", ErrInvalidString
		}
		if !unicode.IsDigit(c) {
			if i+2 < len(r) && unicode.IsDigit(r[i+1]) && unicode.IsDigit(r[i+2]) {
				return "", ErrInvalidString
			}
			count := 1
			if i+1 < len(r) {
				cnt, convertFailed := strconv.Atoi(string(r[i+1]))
				if convertFailed == nil {
					count = cnt
				}
			}
			res += strings.Repeat(string(c), count)
		}
	}
	return res, nil
}
