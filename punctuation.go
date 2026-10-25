package main

import "strings"

func FixPunctuation(words []string) []string {
	final := []string{}
	for i := 0; i < len(words); i++ {
		mot := words[i]
		p := 0
		for p < len(mot) && strings.Contains(".,!?:;", string(mot[p])) { // je compte les signes au debut du mot
			p++
		}
		if p > 0 && len(final) > 0 {
			final[len(final)-1] = final[len(final)-1] + mot[:p] // je colle les signes au mot davant
			if p < len(mot) {
				final = append(final, mot[p:])
			}
		} else {
			final = append(final, mot)
		}
	}
	return final
}
