package strutil

import (
	"strings"
	"unicode"
)

func InsensitiveCmp(str1, str2 string) bool {
	return strings.ToLower(str1) == strings.ToLower(str2)
}

func ToCamelCase(str string) string {
	var (
		sb      strings.Builder
		toUpper bool = true
	)
	for i := 0; i < len(str); i++ {
		//skip the character if it is not a letter
		if !unicode.IsLetter(rune(str[i])) {
			toUpper = true
			continue
		}
		if toUpper {
			sb.WriteRune(unicode.ToUpper(rune(str[i])))
			toUpper = false
		} else {
			sb.WriteByte(str[i])
		}
	}
	return sb.String()
}
