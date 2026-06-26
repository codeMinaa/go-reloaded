package main

import (
	"regexp"
)

func Quote(s string) string {

	//  cleanse spaces
	rea := regexp.MustCompile(`\s+`)
	s = rea.ReplaceAllString(s, " ")

	//  fix single quotes
	re := regexp.MustCompile(`'\s*(.*?)\s*'`)
	s = re.ReplaceAllString(s, "'$1'")

	//  fix double quotes
	re1 := regexp.MustCompile(`"\s*(.*?)\s*"`)
	s = re1.ReplaceAllString(s, `"$1"`)

	return s
}
