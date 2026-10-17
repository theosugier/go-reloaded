package main

import "strings" // pour Contains

func FixPunctuation(words []string) []string { // recoit les mots de result et renvoie les mots avec la ponctuation recollee
	final := []string{}               // liste vide ou on range les mots corriges
	for i := 0; i < len(words); i++ { // on passe sur chaque mot un par un
		mot := words[i]                                                  // le mot actuel
		p := 0                                                           // compteur de signes de ponctuation au debut du mot
		for p < len(mot) && strings.Contains(".,!?:;", string(mot[p])) { // tant quon depasse pas la fin du mot et que la lettre a la position p est un signe
			p++ // on compte un signe de plus
		}
		if p > 0 && len(final) > 0 { // si le mot commence par de la ponctuation et quil y a un mot avant
			final[len(final)-1] = final[len(final)-1] + mot[:p] // on colle la ponctuation a la fin du dernier mot de final genre there + virgule
			if p < len(mot) {                                   // si il reste des lettres apres la ponctuation genre and dans virgule and
				final = append(final, mot[p:]) // on ajoute le reste du mot dans final
			}
		} else { // sinon cest un mot normal
			final = append(final, mot) // on lajoute tel quel
		}
	}
	return final // on renvoie les mots corriges a ProcessText
}
