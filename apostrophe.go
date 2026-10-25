package main

func FixApostrophe(words []string) []string {
	final := []string{}
	ouvert := true // true le prochain ouvre false il ferme
	for i := 0; i < len(words); i++ {
		mot := words[i]
		if words[i] == "'" && ouvert == true && i+1 < len(words) {
			words[i+1] = "'" + words[i+1] // je colle le guillemet devant le mot dapres
			ouvert = false
		} else if words[i] == "'" && ouvert == false && len(final) > 0 {
			final[len(final)-1] = final[len(final)-1] + "'" // je colle le guillemet derriere le mot davant
			ouvert = true
		} else {
			final = append(final, mot)
		}
	}
	return final
}
