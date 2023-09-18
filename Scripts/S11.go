// Escreva um programa que receba uma string e remova todas as vogais. Informe ao usuário o resultado.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Escreva uma string: ")
	scanner.Scan()
	s1 := scanner.Text()

	r := strings.ReplaceAll(s1, "a", "")
	r = strings.ReplaceAll(r, "e", "")
	r = strings.ReplaceAll(r, "i", "")
	r = strings.ReplaceAll(r, "o", "")
	r = strings.ReplaceAll(r, "u", "")

	fmt.Print("Sua string sem vogais fica: ", r)
}
