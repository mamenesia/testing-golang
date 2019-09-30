package main

import (
	"regexp"
	"strings"
)

func getCount(s string) (int, int) {
	var tempVowel string
	var tempConsonant string

	s = strings.ToLower(s)
	splitStr := strings.Split(s, "")
	vowelChar := regexp.MustCompile(`[aeiou]`)

	for i := range splitStr {
		if vowelChar.MatchString(splitStr[i]) {
			if !strings.ContainsAny(tempVowel, splitStr[i]) {
				tempVowel += splitStr[i]
			} else {
				tempVowel = tempVowel
			}
		} else {
			if !strings.ContainsAny(tempConsonant, splitStr[i]) {
				tempConsonant += splitStr[i]
			} else {
				tempConsonant = tempConsonant
			}
		}
	}

	return len(tempVowel), len(tempConsonant)
}
