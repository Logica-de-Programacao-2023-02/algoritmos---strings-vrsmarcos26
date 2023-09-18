//Escreva um programa que receba uma string e substitua todas as ocorrências
//desse caractere na string por outro caractere especificado pelo usuário.
//Informe ao usuário o resultado.

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
	p1 := scanner.Text()
	var nc, vc string
	fmt.Print("Escreva qual caracter você quer que mude na sua string: ")
	fmt.Scan(&vc)
	fmt.Print("Escreva qual caracter você quer que fique no lugar do novo: ")
	fmt.Scan(&nc)

	r := strings.ReplaceAll(p1, vc, nc)

	fmt.Print("Assim fica sua string: ", r)
}
