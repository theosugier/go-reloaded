package main

import (
	"strconv"
	"strings"
)

func ProcessText(text string) string {
	result := []string{}
	words := strings.Fields(text) // je coupe le text en mots / modif strings.Fields pour gérer si il y a plsueiurs espace et on retire argu il prend que 1 argu

	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" {
			nombre, err := strconv.ParseInt(result[len(result)-1], 16, 64) // je lis le mot davant en hexa
			if err == nil {
				result[len(result)-1] = strconv.FormatInt(nombre, 10)
			} // si le mot davant cest pas de lhexa je le laisse sinon il devient 0
		} else if words[i] == "(bin)" {
			nombre, err := strconv.ParseInt(result[len(result)-1], 2, 64)
			if err == nil { // pareil que hex
				result[len(result)-1] = strconv.FormatInt(nombre, 10)
			}
		} else if words[i] == "(up)" {
			result[len(result)-1] = strings.ToUpper(result[len(result)-1])
		} else if words[i] == "(low)" {
			result[len(result)-1] = strings.ToLower(result[len(result)-1])
		} else if words[i] == "(cap)" {
			mot := result[len(result)-1]
			result[len(result)-1] = strings.ToUpper(mot[:1]) + strings.ToLower(mot[1:]) // premiere lettre en maj le reste en min
		} else if words[i] == "(up," && i+1 < len(words) { // verif ya bien un mot apers up
			n, _ := strconv.Atoi(strings.TrimSuffix(words[i+1], ")")) // je recup le nombre dans le mot dapres genre 2)
			if n > len(result) {
				n = len(result) // empeche plant si nombre trop grand
			}
			for j := 1; j <= n; j++ {
				result[len(result)-j] = strings.ToUpper(result[len(result)-j])
			}
			i++ // je saute le 2) vu que je lai deja utilise
		} else if words[i] == "(low," && i+1 < len(words) {
			n, _ := strconv.Atoi(strings.TrimSuffix(words[i+1], ")"))
			if n > len(result) {
				n = len(result) // empeche plant si nombre trop grand
			}
			for j := 1; j <= n; j++ {
				result[len(result)-j] = strings.ToLower(result[len(result)-j])
			}
			i++
		} else if words[i] == "(cap," && i+1 < len(words) {
			n, _ := strconv.Atoi(strings.TrimSuffix(words[i+1], ")"))
			if n > len(result) {
				n = len(result) // empeche plant si nombre trop grand
			}
			for j := 1; j <= n; j++ {
				mot := result[len(result)-j]
				result[len(result)-j] = strings.ToUpper(mot[:1]) + strings.ToLower(mot[1:])
			}
			i++
		} else {
			result = append(result, words[i])
		}
	}
	result = FixPunctuation(result) // apres les marqueurs je fais la ponctuation les apostrophes et a an
	result = FixApostrophe(result)
	result = FixArticle(result)
	return strings.Join(result, " ")
}
