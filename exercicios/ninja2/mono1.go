/*- Escreva um programa que mostre um número em decimal, binário, e hexadecimal. */
package main

import (
	"fmt"
)

func main() {
	a := 1000

	fmt.Printf("decimal = %d - binário = %b - hexadecimal = %#X", a, a, a)
}