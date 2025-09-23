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

	fmt.Scanf("%d", &comando)

	fmt.Println("O endereco da minha variavel comando é:", &comando)
	fmt.Println("O comando escolhido foi:", comando)

}
