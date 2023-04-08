package main

import (
	"fmt"
	"strconv"
)

func isCrescent(s string) bool {
	anterior := -1
	for _, c := range s {
		numero, err := strconv.Atoi(string(c))
		if err != nil {
			return false
		}
		if numero <= anterior {
			return false
		}
		anterior = numero
	}
	return true
}

func main() {
	var s string
	fmt.Print("Diga uma string de números: ")
	fmt.Scan(&s)

	if isCrescent(s) {
		fmt.Println("Os números estão em ordem crescente.")
	} else {
		fmt.Println("Os números não estão em ordem crescente.")
	}
}