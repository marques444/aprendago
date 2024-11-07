/*- Crie uma variável de tipo string utilizando uma raw string literal.
- Demonstre-a.*/
package main

import (
	"fmt"
)

func main() {
	a := `eu 
		odeio
			treinar perna😢`
	fmt.Printf("%v", a)

}