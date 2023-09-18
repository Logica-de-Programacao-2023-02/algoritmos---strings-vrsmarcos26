// Escreva um programa que receba uma string e inverta a ordem dos caracteres. Informe ao usuário o resultado.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Informe uma string: ")
	scanner.Scan()
	s1 := scanner.Text()

	var si string
	for i := len(s1) - 1; i >= 0; i-- {
		si += string(s1[i])
	}

	fmt.Print("Sua string invertida fica: ", si)
}
