package utfdecode

import (
	"html"
	"regexp"
	"strconv"
	"strings"
)

var m = regexp.MustCompile("\\\\u[0-9A-Fa-f]{4,}")

// Decode takes any string with escaped UTF-32 / UTF-16 and prints it out
// as un-escaped (decoded) UTF-16.
func Decode(str string) string {
	for _, s := range strings.Fields(str) {
		us := m.FindAllString(s, -1)
		for i := 0; i < len(us); i++ {
			if isSurrogate(us[i]) {
				e := convertToUTF16(us[i], us[i+1])
				str = strings.Replace(str, us[i], e, 1)
				str = strings.Replace(str, us[i+1], "", 1)
				i++
			} else {
				e := html.UnescapeString("&#x" + strings.ToLower(us[i][2:]) + ";")
				str = strings.Replace(str, us[i], e, 1)
			}
		}
	}
	return str
}

func isSurrogate(s1 string) bool {
	s1 = strings.TrimPrefix(strings.ToLower(s1), `\u`)
	i, _ := strconv.ParseInt("0x"+s1, 0, 64)
	return i >= 0xD800 && i <= 0xDB7F
}

// https://en.wikipedia.org/wiki/UTF-16
func convertToUTF16(s1, s2 string) string {
	s1 = strings.TrimPrefix(strings.ToLower(s1), `\u`)
	s2 = strings.TrimPrefix(strings.ToLower(s2), `\u`)
	i, _ := strconv.ParseInt("0x"+s1, 0, 64)
	j, _ := strconv.ParseInt("0x"+s2, 0, 64)
	a := (i - 0xD800) * 0x400
	b := j - 0xDC00
	c := a + b + 0x10000
	if c < 0 || a < 0 || b < 0 || len(s1) > 4 || len(s2) > 4 {
		return html.UnescapeString("&#x"+s1+";") + html.UnescapeString("&#x"+s2+";")
	}
	str := html.UnescapeString("&#" + strconv.Itoa(int(c)) + ";")
	return str
}
