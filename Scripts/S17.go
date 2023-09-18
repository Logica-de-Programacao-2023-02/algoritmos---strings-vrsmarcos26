// Solicite ao usuário uma string e imprima somente as suas letras únicas (que aparecem apenas uma vez)
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Fale uma string: ")
	scanner.Scan()
	s1 := scanner.Text()

	f := make(map[rune]int)

	for _, caractere := range s1 {
		f[caractere]++
	}

	fmt.Print("Letras únicas na string: ")
	for _, caractere := range s1 {
		if f[caractere] == 1 {
			fmt.Printf("%c ", caractere)
		}
	}

}
