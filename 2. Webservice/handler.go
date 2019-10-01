package main

import (
	"encoding/json"
	"net/http"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
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
		result := []struct {
			Status    int
			Message   string
			Vowel     int
			Consonant int
		}{
			{200, "Successfully counting vowels and consonants.", countVowel, countConsonant},
		}

		resultInBytes, err := json.Marshal(result)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(resultInBytes)

	default:
		http.Error(w, "", http.StatusBadRequest)
	}
}
