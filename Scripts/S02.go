// Escreva um programa que receba uma string e remova todos os espaços em branco.
// Informe ao usuário o resultado.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Escreva uma frase com espaços: ")
	scanner.Scan()
	f1 := scanner.Text()

	s := strings.ReplaceAll(f1, " ", "")

	fmt.Println("Sua frase sem espaços fica: ", s)
}
