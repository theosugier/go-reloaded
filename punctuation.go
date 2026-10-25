package main

import "strings"

func FixPunctuation(words []string) []string {
	final := []string{}
	for i := 0; i < len(words); i++ {
		mot := words[i]
		p := 0
		for p < len(mot) && strings.Contains(".,!?:;", string(mot[p])) {
			p++
		}
		if p > 0 && len(final) > 0 {
			final[len(final)-1] = final[len(final)-1] + mot[:p]
			if p < len(mot) {
				final = append(final, mot[p:])
			}
		} else {
			final = append(final, mot)
		}
	}
	return final
}
