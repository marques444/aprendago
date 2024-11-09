/*- Crie constantes tipadas e não-tipadas.
- Demonstre seus valores.*/
package main

import (
	"fmt"
)

func main() {
	const a = 10
	const b float64 = 10.10
	fmt.Printf("%v-%v", a, b)
}