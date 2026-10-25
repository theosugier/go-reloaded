package main

import "strings"

func FixArticle(words []string) []string {
	for i := 0; i < len(words); i++ {
		if (words[i] == "a" || words[i] == "A") && i+1 < len(words) && strings.Contains("aeiouhAEIOUH", string(words[i+1][0])) {
			words[i] = words[i] + "n"
		}
	}
	return words
}
