package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Print("Diga uma string (sem espaços): ")
	fmt.Scan(&s)

	_, err := strconv.Atoi(s)

	if err != nil {
		fmt.Println("A string não contém só números.")
	} else {
		fmt.Println("A string contém só números.")
	}
}