package main

import "testing"

func TestGetCount(t *testing.T) {
	// Test with input "omama"
	v, c := getCount("omama")

	if v != 2 {
		t.Errorf("Huruf hidup di kata 'omama' ada 2, tapi output %v", v)
	}

	if c != 1 {
		t.Errorf("Huruf mati di kata 'omama' ada 2, tapi output %v", c)
	}

	// Test with input "Dedy Prasetyo H"
	v, c = getCount("Dedy Prasetyo H")

	if v != 3 {
		t.Errorf("Huruf hidup di kata 'Dedy Prasetyo H' ada 3, tapi output %v", v)
	}

	if c != 8 {
		t.Errorf("Huruf hidup di kata 'Dedy Prasetyo H' ada 8, tapi output %v", v)
	}

}
