/*- Utilizando a solução anterior, adicione as opções else if e else.
- Solução: https://play.golang.org/p/_cO7E-yL0o*/
package main

import(
	"fmt"
)

func main(){
	idade:=18
	if idade==18 {
		fmt.Println("Parabéns ce tem 18")
	}else if idade<18{
		fmt.Println("ce é de menor garai")
	}else{
		fmt.Println("ce é de maior ja em")
	}
}