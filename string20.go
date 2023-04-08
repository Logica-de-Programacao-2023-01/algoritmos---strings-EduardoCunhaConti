package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	fmt.Print("Diga uma string (sem espaços): ")
	fmt.Scan(&s)

	camelcase := false
	for _, c := range s {
		if unicode.IsUpper(c) {
			camelcase = true
			break
		}
	}
	if camelcase {
		fmt.Println("A string está em CamelCase.")
	} else {
		fmt.Println("A string não está em CamelCase.")
	}

	numerodepalavras := 0
	for _, c := range s {
		if unicode.IsUpper(c) {
			numerodepalavras++
		}
	}
	fmt.Printf("A string contém %d palavras.\n", numerodepalavras)
}
