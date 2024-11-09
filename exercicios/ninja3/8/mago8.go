/*- Crie um programa que utilize a declaração switch, sem switch statement especificado.
- Solução: https://play.golang.org/p/TyIGp4Hi8B */
package main

import(
	"fmt"
)

func main(){
	idade:=10

	switch{
		case idade==18 :
				fmt.Println("Parabéns ce tem 18")
		case idade<18:
				fmt.Println("ce é de menor garai")
		case idade>18:
				fmt.Println("ce é de maior ja em")
	}			
}