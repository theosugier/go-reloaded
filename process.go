package main

import (
	"strconv" // pour convertir text en nombre et nombre en text avec ParseInt et FormatInt
	"strings" // pour Split Join ToUpper et ToLower
)

func ProcessText(text string) string { // crée la func qui va modifier le text prend text ( un string ) et renvoie un string modifié donc prend sample et renvoie result
	result := []string{} // on cree une liste vide de mots cest la quon met les mots du text final
	words := strings.Split(text, " ") // on coupe le text a chaque espace et on met chaque mot dans words

	for i := range words { // on passe sur chaque mot un par un i cest la position du mot dans words
		if words[i] == "(hex)" { // si le mot cest hex alors le mot davant est un nombre en hexa
			nombre, _ := strconv.ParseInt(result[len(result)-1], 16, 64) // on prend le dernier mot de result on le lis en base 16 et on met le nombre dans nombre le _ cest lerreur quon ignore
			result[len(result)-1] = strconv.FormatInt(nombre, 10) // on remet nombre en text en base 10 donc en decimal a la place du dernier mot
		} else if words[i] == "(bin)" { // pareil mais le mot davant est en binaire
			nombre, _ := strconv.ParseInt(result[len(result)-1], 2, 64) // pareil que hex mais on lis en base 2
			result[len(result)-1] = strconv.FormatInt(nombre, 10) // on remet en decimal a la place du dernier mot
		} else if words[i] == "(up)" { // si cest up on met le mot davant en majuscule
			result[len(result)-1] = strings.ToUpper(result[len(result)-1]) // le dernier mot devient le dernier mot en maj
		} else if words[i] == "(low)" { // pareil que up mais en minuscule
			result[len(result)-1] = strings.ToLower(result[len(result)-1]) // le dernier mot devient le dernier mot en min
		} else if words[i] == "(cap)" { // si cest cap premiere lettre en maj et le reste en min
			mot := result[len(result)-1] // on met le dernier mot dans mot pour pas reecrire result[len(result)-1] a chaque fois
			result[len(result)-1] = strings.ToUpper(mot[:1]) + strings.ToLower(mot[1:]) // mot[:1] cest la premiere lettre quon met en maj mot[1:] cest le reste en min et le + colle les 2
		} else { // si cest pas un marqueur alors cest un mot normal
			result = append(result, words[i]) // on ajoute le mot a la fin de result les marqueurs on les ajoute jamais donc ils disparaissent
		}
	}

	return strings.Join(result, " ") // on recolle tous les mots de result avec un espace entre et on renvoie le text final a main
}
