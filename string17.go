package main

import "fmt"

func main() {
	var s string
	fmt.Println("Diga uma string (sem espaços): ")
	fmt.Scan(&s)

	frequencia := make(map[rune]int)
	for _, c := range s {
		frequencia[c]++
	}

	fmt.Println("Caracteres que aparecem somente uma vez:")
	for c, quantidade := range frequencia {
		if quantidade == 1 {
			fmt.Printf("%c", c)
		}
	}
}