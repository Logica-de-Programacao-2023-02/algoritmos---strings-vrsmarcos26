// Escreva um programa que receba uma string e verifique se ela é um palíndromo. Informe ao usuário se é ou não
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
	if si == s1 {
		fmt.Printf("A string %s é um palíndromo. %s", s1, si)
	} else {
		fmt.Print("Essa string não é um palíndromo.")
	}
}
