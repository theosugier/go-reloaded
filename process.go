package main

import (
  	"strconv"
  	"strings"
  )

func ProcessText(text string) string {
  	result := []string{}
  	words := strings.Split(text, " ")

  	for i := range words {
      		if words[i] == "(hex)" {
            			nombre, _ := strconv.ParseInt(result[len(result)-1], 16, 64)
            			result[len(result)-1] = strconv.FormatInt(nombre, 10)
            		} else {
            			result = append(result, words[i])
            		}
      	}

  	return strings.Join(result, " ")
  }
