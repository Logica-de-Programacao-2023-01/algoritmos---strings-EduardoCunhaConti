package main

import "fmt"

func reverso(s string) string {

	caracteres := []rune(s)

	for i, j := 0, len(caracteres)-1; i < j; i, j = i+1, j-1 {
		caracteres[i], caracteres[j] = caracteres[j], caracteres[i]
	}
	return string(caracteres)
}

func main() {
	var s string

	fmt.Print("Diga uma string (sem espaços e minúscula): ")
	fmt.Scan(&s)

	reverso := reverso(s)

	if s == reverso {
		fmt.Println("A string é um palíndromo.")
	} else {
		fmt.Println("A string não é um palíndromo.")
	}
}
