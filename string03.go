package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	frase := scanner.Text()

	var caractere1 string
	fmt.Print("Diga um caractere da string: ")
	fmt.Scan(&caractere1)

	var caractere2 string
	fmt.Print("Diga um novo caractere: ")
	fmt.Scan(&caractere2)

	nova := strings.ReplaceAll(frase, caractere1, caractere2)
	fmt.Println(nova)
}
