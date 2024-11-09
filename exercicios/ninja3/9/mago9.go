/*- Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".
- Solução: https://play.golang.org/p/4-iTPZwfEz */
package main

import(
	"fmt"
)

func main(){
	esporteFavorito:= "futebol"

	switch esporteFavorito{
		case "futebol":
			fmt.Println("Seu esporte favorito é Futebol!")
		case "basquete":
			fmt.Println("Seu esporte favorito é Basquete!")
		case "volei":
			fmt.Println("Seu esporte favorito é Volei!")
	}
}