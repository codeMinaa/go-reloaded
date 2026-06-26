package main

import (
	"strings"
)

func Articles(text string) string {
	s := strings.Fields(text)
	consonants := "bcdfgjklmnpqrstvwxyzBCDFGJKLMNPQRSTVWXYZ"
	for i := 0; i < len(s); i++ {
		if s[i] == "a" || s[i] == "A" {
			if i+1 < len(s) {
				next := s[i+1]
				first := rune(next[0])
				if strings.ContainsRune("aeiouhAEIOUH", first) {
					if s[i] == "a" {
						s[i] = "an"
					} else {
						s[i] = "An"
					}
				}
			}
		} else if s[i] == "an" || s[i] == "An" {
			if i+1 < len(s) {
				next := s[i+1]
				first := rune(next[0])
				if strings.ContainsRune(consonants, first) {
					if s[i] == "an" {
						s[i] = "a"
					} else {
						s[i] = "A"
					}
				}
			}

		}
	}
	return strings.Join(s, " ")
}
