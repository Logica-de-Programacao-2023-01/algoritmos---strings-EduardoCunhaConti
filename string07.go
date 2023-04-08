package main

import (
	"fmt"
	"regexp"
)

func main() {
	var s string

	fmt.Print("Digite uma string (sem espaços): ")
	fmt.Scan(&s)

	numero := regexp.MustCompile("[0-9]+")
	encontrado := numero.FindStringIndex(s)

	if encontrado == nil {
		fmt.Println("A string inserida não contém um número.")
	} else {
		fmt.Printf("A string inserida contém um número na posição %d\n", encontrado[0])
	}
}
