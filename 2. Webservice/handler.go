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
