package main

import "testing"

func TestSortByVowel(t *testing.T) {
	// Test with input "omama"
	s := sortByVowel("omama")

	if s != "aaomm" {
		t.Errorf("Hasil sort dari kata 'omama' adalah 'aaomm', tapi output %s", s)
	}

	// Test with input "osama"
	s = sortByVowel("osama")

	if s != "aaoms" {
		t.Errorf("Hasil sort dari kata 'osama' adalah 'aaoms', tapi output %s", s)
	}

	// Test with input "patrick"
	s = sortByVowel("patrick")

	if s != "aickprt" {
		t.Errorf("Hasil sort dari kata 'patrick' adalah 'aickprt', tapi output %s", s)
	}
}
