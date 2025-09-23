package main

import "fmt"

func main() {

	leComent()
	comando1 := ledados()

	switch comando1 {
	case "STEVEN":
		fmt.Println("Feraaaaaa")
	case "EDUARDO":
		fmt.Println("Feraaaaaaaaaaaaaa")
	case "MARCOS":
		fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAA")
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
