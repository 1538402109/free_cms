package util

import (
	"fmt"
	"strconv"
	"strings"
)

func Unicode(str string) (textUnquoted string) {
	textQuoted := strconv.QuoteToASCII(str)
	textUnquoted = textQuoted[1 : len(textQuoted)-1]
	return
}

func Ununicode(textUnquoted string) (context string) {
	sUnicodev := strings.Split(textUnquoted, "\\u")
	for _, v := range sUnicodev {
		if len(v) < 1 {
			continue
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			panic(err)
		}
		context += fmt.Sprintf("%c", temp)
	}
	return
}