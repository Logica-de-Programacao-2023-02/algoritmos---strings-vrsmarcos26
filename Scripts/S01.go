//Escreva um programa que receba uma string e converta todas as letras
//minúsculas para maiúsculas. Informe ao usuário o resultado.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Escreva string apenas com letras minúsculas: ")
	scanner.Scan()
	s := scanner.Text()

	r := strings.ToUpper(s)
	fmt.Println("Essa mesma palavra toda maiúscola: ", r)
}
