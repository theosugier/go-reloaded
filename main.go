package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 { // verifie si dans la commande il y a bien 3 arg ( main.go, sample.txt et result.txt)
		fmt.Println("Usage: go-reloaded input output") // si il y a pas renvoie result
		os.Exit(1)                                     // sort du if avec exit on stop tout
	}
	content, err := os.ReadFile(os.Args[1]) // on crée une var content et err et os.ReadFile lis se quil y a dans premiere arg1 ( sample.txt )
	// on met le contenu de args1 dans content avec read et lerreur dans err
	if err != nil { // si err est pas vide alors on print lerr
		fmt.Println("Erreur:", err) // print err
		os.Exit(1)                  // sort du if on stop tout
	}

	result := ProcessText(string(content)) // convertis les bytes en text  (string(content) envoie a ProcessText et attribue result a la var result
	// car Process attend une string pas des bytes

	err = os.WriteFile(os.Args[2], []byte(result), 0644) // on prend se quil y a dans content et on print dans args2 ( result.txt ) et 0644 permission de base par default
	// pareil en haut Write attend des bytes pas du text donc on reconvertis en bytes le text convertis au dessus
	if err != nil {                               // si cela fais une err alors on print
		fmt.Println("Erreur:", err) // print err
		os.Exit(1)                  // on sort on stop
	}
}
