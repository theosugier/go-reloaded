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
	if result != "GO" {
		t.Errorf("erreur up")
	}
	result = ProcessText("HELLO (low)")
	if result != "hello" {
		t.Errorf("erreur low")
	}
	result = ProcessText("bRIDGE (cap)")
	if result != "Bridge" {
		t.Errorf("erreur cap")
	}
	result = ProcessText("this is so exciting (up, 2)")
	if result != "this is SO EXCITING" {
		t.Errorf("erreur up 2")
	}
}