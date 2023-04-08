package main

import "fmt"

func inverso(s string) string {

	caracteres := []rune(s)

	for i, j := 0, len(caracteres)-1; i < j; i, j = i+1, j-1 {
		caracteres[i], caracteres[j] = caracteres[j], caracteres[i]
	}
	return string(caracteres)
}

func main() {
	var s string

	fmt.Print("Diga uma string (sem espaços): ")
	fmt.Scan(&s)

	inverso := inverso(s)
	fmt.Println(inverso)
}
