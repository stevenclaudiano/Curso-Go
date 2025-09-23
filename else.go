package main

import "fmt"

func main() {

	nome := "Steven"
	versao := 1.0
	fmt.Println("OLA EU SOU O:", nome)
	fmt.Println("O programa esta na versão:", versao)

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - exibir log")
	fmt.Println("0 - Sair do Programa")

	var comando int

	fmt.Scan(&comando)

	fmt.Println("O comando escolhido foi:", comando)

	if comando == 1 {

		fmt.Println("Monitorando log.....")

	} else if comando == 2 {

		fmt.Println("Exibindo Log......")

	} else if comando == 0 {

		fmt.Println("Saindo do programa.....")

	} else {

		fmt.Println("Não conheco este comando")
	}

}
