package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Input kata/kalimat yang ingin di cek huruf hidup/mati : ")
	input, _ = reader.ReadString('\n')

	countVowel, countConsonant := getCount(input)
	fmt.Printf("huruf mati: %v, huruf hidup: %v", countVowel, countConsonant)
}
