package main

import "testing"

func TestProcessText(t *testing.T) {

	result := ProcessText("1E (hex)")
	if result != "30" {
		t.Errorf("erreur hex")
	}
	result = ProcessText("10 (bin)")
	if result != "2" {
		t.Errorf("erreur bin")
	}
	result = ProcessText("go (up)")
	if result != "2" {
		t.Errorf("erreur up")
	}
}
