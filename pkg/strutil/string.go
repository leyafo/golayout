package strutil

import(
	"strings"
)

func InsensitiveCmp(str1, str2 string)bool{
	return strings.ToLower(str1) == strings.ToLower(str2)
}