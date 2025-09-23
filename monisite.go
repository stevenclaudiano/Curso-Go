package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibIntroducao()
	exibMenu()
	comando := lecomando()

	switch comando {
	case 1:
		iniciarMonitoamento()
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}

}

func devolveNome() string {
	nome := "Steven"
	return nome
}

func exibIntroducao() {

	nome := "Steven"
	versao := 1.0
	fmt.Println("OLA EU SOU O:", nome)
	fmt.Println("O programa esta na versão:", versao)

}

func exibMenu() {

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - exibir log")
	fmt.Println("0 - Sair do Programa")

}

func lecomando() int {

	var comandolido int
	fmt.Scan(&comandolido)
	fmt.Println("O comando escolhido foi:", comandolido)

	return comandolido

}

func iniciarMonitoamento() {
	fmt.Println("Monitoramento...")
	site := "https://www.alura.com.br/"
	resp, _ := http.Get(site)
	//fmt.Println(resp)

	if resp.StatusCode == 200 {
		fmt.Println("CARREGAMENTO CONCLUIDO!!", site)
	} else {
		fmt.Println("SITE NÃO CARREGADO", site, "E", resp.StatusCode)
	}
}
