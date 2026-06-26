package main

import (
	"regexp"
)

func punctuation(s string) string {

	// Step 2: remove space before punctuation
	rePunc := regexp.MustCompile(`\s+([,.!;:?])`)
	s = rePunc.ReplaceAllString(s, "$1")

	// Step 3: merge repeated punctuation
	reMulti := regexp.MustCompile(`([,.!;:?])([a-zA-Z0-9])`)
	s = reMulti.ReplaceAllString(s, "$1 $2")

	// Step 1: normalize spaces
	reSpaces := regexp.MustCompile(`\s+`)
	s = reSpaces.ReplaceAllString(s, " ")

	return s
}
