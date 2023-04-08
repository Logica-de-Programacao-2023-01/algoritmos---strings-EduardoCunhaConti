package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Println("Diga uma string de números: ")
	fmt.Scan(&s)

	var numeros []int
	for _, c := range s {
		numero, err := strconv.Atoi(string(c))
		if err != nil {
			fmt.Println("String inválida.")
			return
		}

		numeros = append(numeros, numero)
	}

	for i := 1; i < len(numeros); i++ {
		if numeros[i] >= numeros[i-1] {
			fmt.Println("Os números não estão em ordem decrescente.")
			return
		}
	}

	fmt.Println("Os números estão em ordem decrescente.")
}
