package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"strings"
)

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/count", handlerCount)
	http.HandleFunc("/sort", handleSort)

	address := ":8080"
	fmt.Printf("Server started at %s\n", address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		data := []struct {
			Status  int
			Message string
		}{
			{200, "Selamat datang di layanan sorting dan penghitung huruf hidup dan huruf mati."},
		}

		dataInBytes, err := json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(dataInBytes)

	default:
		http.Error(w, "", http.StatusBadRequest)
	}
}

func handlerCount(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		data := []struct {
			Status  int
			Message string
		}{
			{200, "Selamat datang di layanan penghitung huruf hidup dan huruf mati. Silahkan input kata dengan menggunakan metode 'POST"},
		}

		dataInBytes, err := json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(dataInBytes)

	case "POST":
		input := r.FormValue("input")
		countVowel, countConsonant := getCount(input)
		data := []struct {
			Status    int
			Message   string
			Vowel     int
			Consonant int
		}{
			{200, "Berhasil menghitung huruf hidup dan huruf mati.", countVowel, countConsonant},
		}

		dataInBytes, err := json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(dataInBytes)

	default:
		http.Error(w, "", http.StatusBadRequest)
	}
}

func handleSort(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		data := []struct {
			Status  int
			Message string
		}{
			{200, "Selamat datang di layanan sorting. Silahkan input kata dengan menggunakan metode 'POST"},
		}

		dataInBytes, err := json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(dataInBytes)

	case "POST":
		input := r.FormValue("input")
		result := sortByVowel(input)
		data := []struct {
			Status  int
			Message string
			Hasil   string
		}{
			{200, "Berhasil sorting input data.", result},
		}

		dataInBytes, err := json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(dataInBytes)

	default:
		http.Error(w, "", http.StatusBadRequest)
	}
}

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
	sort.Strings(tempConsonant)
	vowel := strings.Join(tempVowel, "")
	consonant := strings.Join(tempConsonant, "")

	result := vowel + consonant

	return result
}
