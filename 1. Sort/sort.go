package main

import (
	"regexp"
	"sort"
	"strings"
)

func sortByVowel(s string) string {
	var tempVowel []string
	var tempConsonant []string

	splitStr := strings.Split(s, "")
	vowelChar := regexp.MustCompile(`[aeiouAEIOU]`)

	for i := range splitStr {
		if vowelChar.MatchString(splitStr[i]) {
			tempVowel = append(tempVowel, splitStr[i])
		} else {
			tempConsonant = append(tempConsonant, splitStr[i])
		}
	}

	sort.Strings(tempVowel)
	Vowel := strings.Join(tempVowel, "")

	sort.Strings(tempConsonant)
	Consonant := strings.Join(tempConsonant, "")

	return Vowel + Consonant
}
