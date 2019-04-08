package orm

import "strings"

func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	num := len(s)

	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' {
			data = append(data, '_')
		}
		data = append(data, d)
	}

	return strings.ToLower(string(data[:]))
}
