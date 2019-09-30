package main

import (
	"fmt"
)

func main() {
	countVowel, countConsonant := getCount("omama")
	fmt.Printf("huruf hiduo: %d, huruf mati: %d", countVowel, countConsonant)
}
