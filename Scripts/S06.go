// Escreva um programa que receba uma string e conte quantas palavras ela contém. Informe ao usuário o resultado.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Fale uma string: ")
	scanner.Scan()
	f1 := scanner.Text()

	p := strings.Fields(f1)

	np := len(p)

	fmt.Printf("Tem exatamente %d palavras na sua frase: %s ", np, f1)

}
