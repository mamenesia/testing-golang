package main

import "fmt"

func main() {
	countVowel, countConsonant := getCount("omama")
	fmt.Printf("huruf mati: %v, huruf hidup: %v", countVowel, countConsonant)
}
