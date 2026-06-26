package main

import (
	"strconv"
	"strings"
)

func Conversion(str string) string {
	s := strings.Fields(str)

	for i := 0; i < len(s); i++ {

		// hex conversion
		if s[i] == "(hex)" && i > 0 {
			input, err := strconv.ParseInt(s[i-1], 16, 64)
			if err != nil {
				return "error in hex number"
			}

			s[i-1] = strconv.FormatInt(input, 10)
			s = append(s[:i], s[i+1:]...)
			i--
		}

		// bin conversion
		if s[i] == "(bin)" && i > 0 {
			input, err := strconv.ParseInt(s[i-1], 2, 64)
			if err != nil {
				return "error in bin number"
			}

			s[i-1] = strconv.FormatInt(input, 10)
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}

	return strings.Join(s, " ")
}
