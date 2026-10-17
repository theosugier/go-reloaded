package main

import "strings" // pour Contains

func FixArticle(words []string) []string { // recoit les mots et transforme a en an quand il faut
	for i := 0; i < len(words); i++ { // on passe sur chaque mot un par un
		if (words[i] == "a" || words[i] == "A") && i+1 < len(words) && strings.Contains("aeiouhAEIOUH", string(words[i+1][0])) { // si le mot cest a ou A et quil y a un mot apres et que ce mot commence par une voyelle ou un h
			words[i] = words[i] + "n" // on ajoute un n donc a devient an et A devient An
		}
	}
	return words // on renvoie les mots corriges a processtext
}
