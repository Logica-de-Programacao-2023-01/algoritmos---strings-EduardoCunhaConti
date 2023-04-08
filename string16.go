package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		s1 string
		s2 string
	)
	fmt.Print("Diga uma string (sem espaços e minúscula): ")
	fmt.Scan(&s1)
	fmt.Print("Diga outra string (sem espaços e minúscula): ")
	fmt.Scan(&s2)

	if strings.Contains(s1, s2) {
		fmt.Println("A segunda string é uma substring da primeira.")
	} else {
		fmt.Println("A segunda string não é uma substring da primeira.")
	}
}
