/*- Crie um loop utilizando a sintaxe: for condition {}
- Utilize-o para demonstrar os anos desde que você nasceu.
- Solução: https://play.golang.org/p/qnFjiDJzLor*/
package main
import(
	"fmt"
)

func main(){
	anoNascimento:=2007
	anoAtual:= 2024

	for anoNascimento<=anoAtual{
		fmt.Println(anoNascimento)
		anoNascimento++
	}
}
