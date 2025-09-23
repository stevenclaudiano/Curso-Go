package main

import "fmt"

func main() {

	exibMenu()

	var comando string
	var num1 float64
	var num2 float64

	fmt.Scan(&comando)

	if comando == "+" {
		fmt.Println("Digite o primeiro número:")
		fmt.Scan(&num1)
		fmt.Println("Digite o segundo número:")
		fmt.Scan(&num2)

		resultado := num1 + num2
		fmt.Println("O resultado da soma é:", resultado)

	} else if comando == "-" {
		fmt.Println("Digite o primeiro número:")
		fmt.Scan(&num1)
		fmt.Println("Digite o segundo número:")
		fmt.Scan(&num2)

		resultado := num1 - num2
		fmt.Println("O resultado da soma é:", resultado)
	} else if comando == "*" {
		fmt.Println("Digite o primeiro número:")
		fmt.Scan(&num1)
		fmt.Println("Digite o segundo número:")
		fmt.Scan(&num2)

		resultado := num1 * num2
		fmt.Println("O resultado da soma é:", resultado)
	} else if comando == "/" {
		fmt.Println("Digite o primeiro número:")
		fmt.Scan(&num1)
		fmt.Println("Digite o segundo número:")
		fmt.Scan(&num2)

		resultado := num1 / num2
		fmt.Println("O resultado da soma é:", resultado)
	}

}

func exibMenu() {
	fmt.Println("Escolha o tipo de cálculo:")
	fmt.Println("+ - Adição")
	fmt.Println("- - Subtração")
	fmt.Println("* - Multiplicação")
	fmt.Println("/ - Divisão")
}
