package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go-reloaded input output")
		os.Exit(1)
	}
	content, err := os.ReadFile(os.Args[1]) // je lis sample.txt
	if err != nil {
		fmt.Println("Erreur:", err)
		os.Exit(1)
	}

	result := ProcessText(string(content)) // jenvoie le text a process

	err = os.WriteFile(os.Args[2], []byte(result), 0644) // jecris le resultat dans result.txt
	if err != nil {
		fmt.Println("Erreur:", err)
		os.Exit(1)
	}
}
