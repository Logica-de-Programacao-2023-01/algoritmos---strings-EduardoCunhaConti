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

	frase = strings.TrimSpace(frase)
	palavras := strings.Split(frase, " ")
	fmt.Printf("A frase inserida contém %d palavras\n", len(palavras))
}
