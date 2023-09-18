// Solicite ao usuário duas strings e informe se elas são iguais ou diferentes.
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
	s2 := scanner.Text()
	fmt.Print("Fale outra string: ")
	scanner.Scan()
	s1 := scanner.Text()
	if s1 == s2 {
		fmt.Printf("As strings %s e %s são iguais.", s1, s2)
	} else {
		fmt.Printf("As strings %s e %s são diferentes.", s1, s1)
	}
}
