package main

import "fmt"

func main() {
	var s string

	fmt.Print("Diga uma string (sem espaços): ")
	fmt.Scan(&s)

	resultado := ""
	for _, c := range s {
		if c != 'a' && c != 'e' && c != 'i' && c != 'o' && c != 'u' {
			resultado += string(c)
		}
	}
	fmt.Println(resultado)
}