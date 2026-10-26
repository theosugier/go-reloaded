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
	result = ProcessText("there ,and then BAMM !!")
	if result != "there, and then BAMM!!" {
		t.Errorf("erreur ponctuation")
	}
	result = ProcessText("me: ' awesome '")
	if result != "me: 'awesome'" {
		t.Errorf("erreur apostrophe")
	}
	result = ProcessText("A amazing rock")
	if result != "An amazing rock" {
		t.Errorf("erreur article")
	}
	result = ProcessText("(up) hello")
	if result != "(up) hello" {
		t.Errorf("erreur debut")
	}
	result = ProcessText("hi (up, 5)")
	if result != "HI" {
		t.Errorf("erreur nombre trop grand")
	}
	result = ProcessText("bonjour (hex)")
	if result != "bonjour" {
		t.Errorf("erreur hex faux")
	}
	result = ProcessText("hello  world")
	if result != "hello world" {
		t.Errorf("erreur espaces")
	}
}
