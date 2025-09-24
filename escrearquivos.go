package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramento = 3
const delay = 5

func main() {
	exibIntroducao()

	registralog("site-falso", false)

	for {

		exibMenu()
		comando := lecomando()

		switch comando {
		case 1:
			iniciarMonitoramento()
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

// restante do código omitido

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := lesitesdoarquivo()

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")

	}

	fmt.Println("")
}

// restante do código omitido

func testaSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registralog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registralog(site, false)
	}
}

func lesitesdoarquivo() []string {

	var sites []string

	arquivo, err := os.Open("arquivo.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites

}

func registralog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(site + " - online: " + fmt.Sprint(status) + "\n")

}
