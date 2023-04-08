package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		s1 string
		s2 string
	)

	fmt.Print("Diga uma string (minúscula e sem espaços): ")
	fmt.Scan(&s1)
	fmt.Print("Diga outra string (minúscula e sem espaços): ")
	fmt.Scan(&s2)

	rune1 := []rune(s1)
	rune2 := []rune(s2)

	sort.Slice(rune1, func(i, j int) bool { return rune1[i] < rune1[j] })
	sort.Slice(rune2, func(i, j int) bool { return rune2[i] < rune2[j] })

	palavra1 := string(rune1)
	palavra2 := string(rune2)

	if palavra1 == palavra2 {
		fmt.Println("As strings são anagramas.")
	} else {
		fmt.Println("As palavras não são anagramas.")
	}
}