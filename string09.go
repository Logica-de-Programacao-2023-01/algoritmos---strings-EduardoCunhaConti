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

	var letra1 string
	fmt.Print("Diga uma letra da string: ")
	fmt.Scan(&letra1)

	var letra2 string
	fmt.Print("Diga uma nova letra: ")
	fmt.Scan(&letra2)

	nova := strings.ReplaceAll(frase, letra1, letra2)
	fmt.Println(nova)
}
