// Escreva um programa que receba uma string e verifique se ela é um número válido em ponto flutuante.
// Informe ao usuário se é ou não.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Print("Escreva uma string: ")
	scanner.Scan()
	s1 := scanner.Text()
	n, err := strconv.ParseFloat(s1, 64)
	if err == nil {
		fmt.Printf("Seu número %s é um número válido em ponto flutuante. %.2f ", s1, n)
	} else {
		fmt.Printf("Sua string %s não é valido em ponto flutuante.", s1)
	}
}
