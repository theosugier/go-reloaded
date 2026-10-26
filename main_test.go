package main

import "testing"

func TestProcessText(t *testing.T) {

	result := ProcessText("1E (hex)")
	if result != "30" {
		t.Errorf("erreur hex")
	}
}
