package main

import (
	"conversorTemperatura/conversor"
	"fmt"
	"os"
)

var (
	nomeDoUsuario       string
	unidadeOriginal     string
	unidadeConvertida   string
	temperaturaInserida float64
)

func main() {
	fmt.Print("Digite o seu nome: ")
	fmt.Scan(&nomeDoUsuario)

	fmt.Println("Olá,", nomeDoUsuario+"!"+"\n"+
		"Bem-vindo ao Conversor de Temperatura!")

	for {
		fmt.Println("Escolha a unidade de medida da temperatura que você irá inserir:")
		fmt.Println("[1] Celsius")
		fmt.Println("[2] Fahrenheit")
		fmt.Println("[3] Kelvin")
		fmt.Println("[4] Cancelar e sair")

		var escolha int
		_, err := fmt.Scan(&escolha)

		if err != nil {
			fmt.Println("Erro ao ler a entrada. Por favor, digite um número.")
			fmt.Scanln()
			continue
		}

		switch escolha {
		case 1:
			unidadeOriginal = "celsius"
		case 2:
			unidadeOriginal = "fahrenheit"
		case 3:
			unidadeOriginal = "kelvin"
		case 4:
			fmt.Println("Programa encerrado. Até logo," + nomeDoUsuario + "!")
			return
		default:
			fmt.Println("Opção inválida. Por favor, escolha 1, 2, 3 ou 4.")
			continue
		}

		fmt.Printf("Digite a temperatura em %s: ", unidadeOriginal)
		_, err = fmt.Scan(&temperaturaInserida)

		if err != nil {
			fmt.Println("Erro ao ler a temperatura. Certifique-se de inserir um número.")
			fmt.Scanln()
			continue
		}

		for {
			fmt.Println("Agora, escolha a unidade para a qual você deseja converter:")
			fmt.Println("[1] Celsius")
			fmt.Println("[2] Fahrenheit")
			fmt.Println("[3] Kelvin")
			fmt.Println("[4] Voltar para o menu anterior")

			_, err = fmt.Scan(&escolha)

			if err != nil {
				fmt.Println("Erro ao ler a entrada. Por favor, digite um número.")
				fmt.Scanln()
				continue
			}

			switch escolha {
			case 1:
				unidadeConvertida = "celsius"
			case 2:
				unidadeConvertida = "fahrenheit"
			case 3:
				unidadeConvertida = "kelvin"
			case 4:
				break
			default:
				fmt.Println("Opção inválida. Por favor, escolha 1, 2, 3 ou 4.")
				continue
			}

			if escolha == 4 {
				break
			}

			conversor.ConverterTemperatura(unidadeOriginal, unidadeConvertida, temperaturaInserida, os.Stdout)

			fmt.Println("Deseja realizar outra conversão?")
			fmt.Println("[1] Sim")
			fmt.Println("[2] Encerrar programa")

			_, err = fmt.Scan(&escolha)

			if err != nil {
				fmt.Println("Erro ao ler a entrada. Por favor, digite um número.")
				fmt.Scanln()
				return
			}

			if escolha == 2 {
				fmt.Println("Programa encerrado. Até logo,", nomeDoUsuario+"!")
				return
			}
		}
	}
}
