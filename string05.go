package main

import (
	"fmt"
	"strconv"
)

func main() {
	var string string

	fmt.Print("Diga uma string (sem espaços): ")
	fmt.Scan(&string)

	i, err := strconv.ParseFloat(string, 64)
	if err != nil {
		fmt.Printf("O valor %s não é um float\n", string)
	} else {
		fmt.Printf("O valor %f\n é float", i)
	}
}
