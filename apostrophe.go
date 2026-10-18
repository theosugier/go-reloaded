package main

func FixApostrophe(words []string) []string {
	final := []string{}
	for i := 0; i < len(words); i++ {
		mot := words[i] 
		final = append(final, mot)
	}  
	return final 
}