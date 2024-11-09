/*- Crie um loop utilizando a sintaxe: for {}
- Utilize-o para demonstrar os anos desde que você nasceu.
- Solução: https://play.golang.org/p/dIbfdiuw0ms*/
package main

import(
	"fmt"
)

func main(){

	anoNascimento:=2007
	anoAtual:=2024
	for{
		if anoNascimento>anoAtual {
			break
		}
		fmt.Println(anoNascimento)
		anoNascimento++
	}

}