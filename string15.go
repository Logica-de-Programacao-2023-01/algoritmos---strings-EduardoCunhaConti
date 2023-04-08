package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Println("Diga uma string (sem espaços): ")
	fmt.Scan(&s)

	vogais := "aeiouAEIOU"

	nova := strings.Builder{}
	for _, c := range s {
		if strings.ContainsRune(vogais, c) {
			nova.WriteRune('*')
		} else {
			nova.WriteRune(c)
		}
	}
	fmt.Println(nova.String())
}