package utfdecode

import (
	"html"
	"regexp"
	"strconv"
	"strings"
)

var m = regexp.MustCompile("\\\\u[0-9A-Fa-f]{4}")

// Decode takes any string with escaped UTF-32 / UTF-16 and prints it out
// as un-escaped (decoded) UTF-16.
func Decode(str string) string {
	for _, s := range strings.Fields(str) {
		us := m.FindAllString(s, -1)
		for i := 0; i < len(us)-1; i += 2 {
			e := convertToUTF16(us[i], us[i+1])
			// fmt.Println(us[i], us[i+1], e)
			str = strings.Replace(str, us[i], e, 1)
			str = strings.Replace(str, us[i+1], "", 1)
		}
		if len(us)%2 == 1 {
			e := us[len(us)-1]
			str = strings.Replace(str, e, html.UnescapeString("&#"+e[2:]+";"), 1)
		}
	}
	return str
}

// https://en.wikipedia.org/wiki/UTF-16
func convertToUTF16(s1, s2 string) string {
	if len(s1) == 6 {
		s1 = s1[2:]
	}
	if len(s2) == 6 {
		s2 = s2[2:]
	}
	first := "0x" + s1
	second := "0x" + s2
	i, _ := strconv.ParseInt(first, 0, 64)
	j, _ := strconv.ParseInt(second, 0, 64)
	a := (i - 0xD800) * 0x400
	b := j - 0xDC00
	c := a + b + 0x10000
	// fmt.Println(s1, s2, a, b, c)
	if c < 0 {
		return html.UnescapeString("&#x"+s1+";") + html.UnescapeString("&#x"+s2+";")
	}
	str := html.UnescapeString("&#" + strconv.Itoa(int(c)) + ";")
	return str
}
