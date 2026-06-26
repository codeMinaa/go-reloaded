package main

import (
	"strconv"
	"strings"
)

func casing(text string) string {
	v := strings.Fields(text)
	for i := 0; i < len(v); i++ {
		if strings.HasPrefix(v[i], "(low") {
			if v[i] == "(low)" {
				if i-1 >= 0 {
					v[i-1] = strings.ToLower(v[i-1])
				}
				v = append(v[:i], v[i+1:]...)
				i--
			}
		}
		if v[i] == "(low," {
			num := strings.TrimSuffix(v[i+1], ")")
			s, _ := strconv.Atoi(num)
			for d := i - 1; d >= 0 && d >= i-s; d-- {
				v[d] = strings.ToLower(v[d])
			}
			v = append(v[:i], v[i+2:]...)
			i--
		}
		if strings.HasPrefix(v[i], "(up") {
			if v[i] == "(up)" {
				if i-1 >= 0 {
					v[i-1] = strings.ToUpper(v[i-1])
				}
				v = append(v[:i], v[i+1:]...)
				i--
			}
		}
		if v[i] == "(up," {
			num := strings.TrimSuffix(v[i+1], ")")
			s, _ := strconv.Atoi(num)
			for d := i - 1; d >= 0 && d >= i-s; d-- {
				v[d] = strings.ToUpper(v[d])
			}
			v = append(v[:i], v[i+2:]...)
			i--
		}

	}
	return strings.Join(v, " ")
}
