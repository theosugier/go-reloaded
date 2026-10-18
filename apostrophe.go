package main

func FixApostrophe(words []string) []string {
	final := []string{}
	ouvert := true 
	for i := 0; i < len(words); i++ {
		mot := words[i] 
		if words[i] == "'" && ouvert == true && i+1 < len(words) {
			words[i+1] = "'" + words[i+1]
			ouvert = false	
		} else {
		final = append(final, mot) 
		}
	}  
	return final 
}