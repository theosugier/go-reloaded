package main

import (
	"strconv" // pour convertir text en nombre et nombre en text avec ParseInt et FormatInt
	"strings" // pour Split Join ToUpper et ToLower
)

func ProcessText(text string) string { // crée la func qui va modifier le text prend text ( un string ) et renvoie un string modifié donc prend sample et renvoie result
	result := []string{}              // on cree une liste vide de mots cest la quon met les mots du text final
	words := strings.Split(text, " ") // on coupe le text a chaque espace et on met chaque mot dans words

	for i := 0; i < len(words); i++ { // on passe sur chaque mot un par un i cest la position du mot dans words on utilise plus range pour pouvoir sauter un mot avec i++
		if words[i] == "(hex)" { // si le mot cest hex alors le mot davant est un nombre en hexa
			nombre, _ := strconv.ParseInt(result[len(result)-1], 16, 64) // on prend le dernier mot de result on le lis en base 16 et on met le nombre dans nombre le _ cest lerreur quon ignore
			result[len(result)-1] = strconv.FormatInt(nombre, 10)        // on remet nombre en text en base 10 donc en decimal a la place du dernier mot
		} else if words[i] == "(bin)" { // pareil mais le mot davant est en binaire
			nombre, _ := strconv.ParseInt(result[len(result)-1], 2, 64) // pareil que hex mais on lis en base 2
			result[len(result)-1] = strconv.FormatInt(nombre, 10)       // on remet en decimal a la place du dernier mot
		} else if words[i] == "(up)" { // si cest up on met le mot davant en majuscule
			result[len(result)-1] = strings.ToUpper(result[len(result)-1]) // le dernier mot devient le dernier mot en maj
		} else if words[i] == "(low)" { // pareil que up mais en minuscule
			result[len(result)-1] = strings.ToLower(result[len(result)-1]) // le dernier mot devient le dernier mot en min
		} else if words[i] == "(cap)" { // si cest cap premiere lettre en maj et le reste en min
			mot := result[len(result)-1]                                                // on met le dernier mot dans mot pour pas reecrire result[len(result)-1] a chaque fois
			result[len(result)-1] = strings.ToUpper(mot[:1]) + strings.ToLower(mot[1:]) // mot[:1] cest la premiere lettre quon met en maj mot[1:] cest le reste en min et le + colle les 2
		} else if words[i] == "(up," { // split a coupe up 2 en deux mot donc on tombe sur le debut du marqueur et le nombre est dans le mot suivant
			n, _ := strconv.Atoi(strings.TrimSuffix(words[i+1], ")")) // on prend le mot dapres genre 2) on enleve la ) avec TrimSuffix puis Atoi transforme le text 2 en vrai nombre et on le met dans n
			for j := 1; j <= n; j++ {                                 // on tourne n fois j vaut 1 puis 2 etc jusqua n
				result[len(result)-j] = strings.ToUpper(result[len(result)-j]) // au tour j on met en maj le mot a la position len(result)-j donc le dernier puis lavant dernier etc
			}
			i++ // on saute le mot dapres genre 2) vu quon la deja utilise sinon il se ferait ajouter dans result
		} else if words[i] == "(low," { // pareil que up 2 mais en minuscule split a coupe low 2 en deux mot donc le nombre est dans le mot suivant
			n, _ := strconv.Atoi(strings.TrimSuffix(words[i+1], ")")) // pareil que up on prend le mot dapres genre 2) on enleve la ) et Atoi le transforme en vrai nombre dans n
			for j := 1; j <= n; j++ {                                 // on tourne n fois j vaut 1 puis 2 etc jusqua n
				result[len(result)-j] = strings.ToLower(result[len(result)-j]) // au tour j on met en min le mot a la position len(result)-j donc le dernier puis lavant dernier etc
			}
			i++ // on saute le mot dapres genre 2) vu quon la deja utilise sinon il se ferait ajouter dans result
		} else if words[i] == "(cap," { // pareil que up 2 mais en cap split a coupe cap 2 en deux mot donc le nombre est dans le mot suivant
			n, _ := strconv.Atoi(strings.TrimSuffix(words[i+1], ")")) // pareil que up on prend le mot dapres genre 2) on enleve la ) et Atoi le transforme en vrai nombre dans n
			for j := 1; j <= n; j++ {                                 // on tourne n fois j vaut 1 puis 2 etc jusqua n
				mot := result[len(result)-j]                                                // au tour j on met le mot a la position len(result)-j dans mot
				result[len(result)-j] = strings.ToUpper(mot[:1]) + strings.ToLower(mot[1:]) // premiere lettre en maj et le reste en min comme cap simple mais sur le mot du tour j
			}
			i++ // on saute le mot dapres genre 2) vu quon la deja utilise sinon il se ferait ajouter dans result
		} else { // si cest pas un marqueur alors cest un mot normal
			result = append(result, words[i]) // on ajoute le mot a la fin de result les marqueurs on les ajoute jamais donc ils disparaissent
		}
	}
	result = FixPunctuation(result)
	result = FixApostrophe(result)
	result = FixArticle(result)
	return strings.Join(result, " ") // on recolle tous les mots de result avec un espace entre et on renvoie le text final a main
}
