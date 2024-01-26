package conversor

import (
	"fmt"
	"io"
)

// ConverterTemperatura realiza a conversão da temperatura entre as unidades especificadas.
func ConverterTemperatura(unidadeOriginal, unidadeConvertida string, temperatura float64, output io.Writer) {
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

	fmt.Fprintf(output, "%.2f %s é igual a %.2f %s\n", temperatura, unidadeOriginal, temperaturaConvertida, unidadeConvertida)
}

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

func converterDeFahrenheit(temperatura float64, unidadeConvertida string) float64 {
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
