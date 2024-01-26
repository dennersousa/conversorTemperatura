package main

import (
	"fmt"
)

var (
	nomeDoUsuario       string
	unidadeOriginal     string
	unidadeConvertida   string
	temperaturaInserida float64
)

// main é a função principal que inicia a execução do programa.
func main() {
	fmt.Print("Digite o seu nome: ")
	fmt.Scan(&nomeDoUsuario)

	fmt.Println("Olá, ", nomeDoUsuario, "!\n"+
		"Bem-vindo ao Conversor de Temperatura!")

	for {
		// Menu principal para escolher a unidade de medida da temperatura.
		fmt.Println("Escolha a unidade de medida da temperatura que você irá inserir:")
		fmt.Println("[1] Celsius")
		fmt.Println("[2] Fahrenheit")
		fmt.Println("[3] Kelvin")
		fmt.Println("[4] Cancelar e sair")

		var escolha int
		_, err := fmt.Scan(&escolha)

		if err != nil {
			fmt.Println("Erro ao ler a entrada. Por favor, digite um número.")
			fmt.Scanln() // Limpar o buffer de entrada
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
			fmt.Println("Programa encerrado. Até logo, ", nomeDoUsuario, "!")
			return
		default:
			fmt.Println("Opção inválida. Por favor, escolha 1, 2, 3 ou 4.")
			continue
		}

		// Solicitar a temperatura ao usuário.
		fmt.Printf("Digite a temperatura em %s: ", unidadeOriginal)
		_, err = fmt.Scan(&temperaturaInserida)

		if err != nil {
			fmt.Println("Erro ao ler a temperatura. Certifique-se de inserir um número.")
			fmt.Scanln() // Limpar o buffer de entrada
			continue
		}

		for {
			// Menu para escolher a unidade de conversão.
			fmt.Println("Agora, escolha a unidade para a qual você deseja converter:")
			fmt.Println("[1] Celsius")
			fmt.Println("[2] Fahrenheit")
			fmt.Println("[3] Kelvin")
			fmt.Println("[4] Voltar para o menu anterior")

			_, err = fmt.Scan(&escolha)

			if err != nil {
				fmt.Println("Erro ao ler a entrada. Por favor, digite um número.")
				fmt.Scanln() // Limpar o buffer de entrada
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

			// Realizar a conversão e exibir o resultado.
			converterTemperatura(unidadeOriginal, unidadeConvertida, temperaturaInserida)

			// Perguntar ao usuário se deseja realizar outra conversão.
			fmt.Println("Deseja realizar outra conversão?")
			fmt.Println("[1] Sim")
			fmt.Println("[2] Encerrar programa")

			_, err = fmt.Scan(&escolha)

			if err != nil {
				fmt.Println("Erro ao ler a entrada. Por favor, digite um número.")
				fmt.Scanln() // Limpar o buffer de entrada
				return
			}

			if escolha == 2 {
				fmt.Println("Programa encerrado. Até logo, ", nomeDoUsuario, "!")
				return
			}
		}
	}
}

// converterTemperatura realiza a conversão da temperatura entre as unidades especificadas.
func converterTemperatura(unidadeOriginal, unidadeConvertida string, temperatura float64) {
	var temperaturaConvertida float64

	switch unidadeOriginal {
	case "celsius":
		temperaturaConvertida = converterDeCelsius(temperatura, unidadeConvertida)
	case "fahrenheit":
		temperaturaConvertida = converterDeFahrenheit(temperatura, unidadeConvertida)
	case "kelvin":
		temperaturaConvertida = converterDeKelvin(temperatura, unidadeConvertida)
	default:
		fmt.Println("Unidade de temperatura original desconhecida.")
		return
	}

	fmt.Printf("%.2f %s é igual a %.2f %s\n", temperatura, unidadeOriginal, temperaturaConvertida, unidadeConvertida)
}

// converterDeCelsius converte a temperatura de Celsius para a unidade especificada.
func converterDeCelsius(temperatura float64, unidadeConvertida string) float64 {
	switch unidadeConvertida {
	case "celsius":
		return temperatura
	case "fahrenheit":
		return (temperatura * 9 / 5) + 32
	case "kelvin":
		return temperatura + 273.15
	default:
		fmt.Println("Unidade de temperatura para conversão desconhecida.")
		return 0
	}
}

// converterDeFahrenheit converte a temperatura de Fahrenheit para a unidade especificada.
func converterDeFahrenheit(tempxeratura float64, unidadeConvertida string) float64 {
	switch unidadeConvertida {
	case "celsius":
		return (temperatura - 32) * 5 / 9
	case "fahrenheit":
		return temperatura
	case "kelvin":
		return (temperatura + 459.67) * 5 / 9
	default:
		fmt.Println("Unidade de temperatura para conversão desconhecida.")
		return 0
	}
}

// converterDeKelvin converte a temperatura de Kelvin para a unidade especificada.
func converterDeKelvin(temperatura float64, unidadeConvertida string) float64 {
	switch unidadeConvertida {
	case "celsius":
		return temperatura - 273.15
	case "fahrenheit":
		return (temperatura * 9 / 5) - 459.67
	case "kelvin":
		return temperatura
	default:
		fmt.Println("Unidade de temperatura para conversão desconhecida.")
		return 0
	}
}
