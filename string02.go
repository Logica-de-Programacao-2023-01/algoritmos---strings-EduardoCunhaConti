package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	frase := scanner.Text()

	resultado := ""
	for _, c := range frase {
		if c != ' ' {
			resultado += string(c)
		}
	}
	fmt.Println(resultado)
}
