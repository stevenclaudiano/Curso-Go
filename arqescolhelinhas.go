package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramento = 3
const delay = 5

func main() {

	//exibeNomes()
	lesitesdoarquivo()
	exibIntroducao()
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

// restante do código omitido

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	/*sites := []string{"https://random-status-code.herokuapp.com/",
	"https://www.alura.com.br", "https://www.caelum.com.br"}

	*/

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
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}

func lesitesdoarquivo() []string {

	var sites []string

	//arquivo, err := os.Open("arquivo.txt")	//Esta apontando para o arquivo

	//arquivo, err := ioutil.ReadFile("arquivo.txt")	// Ler arquivos de texto

	arquivo, err := os.Open("arquivo.txt") //Esta apontando para o arquivo

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	linha, err := leitor.ReadString('\n')
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(linha)
	//fmt.Println(string(arquivo))
	return sites
}
