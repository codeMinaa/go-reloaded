package main

import (
	"strconv"
	"strings"
)

func Capitalized(text string) string {
	x := strings.Fields(text)
	for i := 0; i < len(x); i++ {
		if strings.HasPrefix(x[i], "(cap") {
			if x[i] == "(cap)" {
				if i-1 >= 0 {
					x[i-1] = strings.ToUpper(x[i-1][:1]) + strings.ToLower(x[i-1][1:])
				}
				x = append(x[:i], x[i+1:]...)
				i--
			}
		}
		if x[i] == "(cap," {
			num := strings.TrimSuffix(x[i+1], ")")
			c, _ := strconv.Atoi(num)
			for z := i - 1; z >= 0 && z >= i-c; z-- {
				x[z] = strings.ToUpper(x[z][:1]) + strings.ToLower(x[z][1:])
			}
			x = append(x[:i], x[i+2:]...)
			i--
		}
	}
	return strings.Join(x, " ")
}
