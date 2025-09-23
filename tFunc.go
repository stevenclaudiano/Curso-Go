package main

import "fmt"

func main() {

	leComent()
	comando1 := ledados()

	switch comando1 {
	case "TESTE1":
		fmt.Println("AAAAAAAAAAAAAAA")
	case "TESTE2":
		fmt.Println("BBBBBBBBBBBBBBBBBBB")
	case "TESTE3":
		fmt.Println("CCCCCCCCCCCCCCCCCCCCCCCC")
	}

}

func leComent() {

	fmt.Println("Escolha uma dessas opções")
	fmt.Println("1 - Status")
	fmt.Println("2 - Programa")
	fmt.Println("3 - Sair do Programa")

}

func ledados() string {

	var comandolido string
	fmt.Scan(&comandolido)
	fmt.Println("O Nome escolhido foi este aqui: ", comandolido)

	return comandolido
}
