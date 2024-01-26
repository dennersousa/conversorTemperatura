# Conversor de Temperatura

Este projeto visa oferecer uma solução eficiente para a conversão de temperaturas entre as unidades Celsius, Fahrenheit e Kelvin, utilizando a linguagem de programação Golang. Desenvolvido com ênfase na confiabilidade do código, o projeto incorpora práticas de teste unitário para garantir a precisão das conversões.

## Objetivos

Os objetivos fundamentais deste projeto são:

1. **Desenvolver um Conversor Eficiente:**
   - Implementar um conversor de temperatura robusto e de fácil utilização.

2. **Garantir Confiabilidade com Testes Unitários:**
   - Utilizar testes unitários para assegurar que o conversor funcione de forma confiável em diversas situações.

3. **Explorar Conceitos Básicos de Golang:**
   - Abordar e aplicar conceitos básicos da linguagem Golang, incluindo entrada e saída, estruturas de dados e funções.

## Abordagem

A abordagem adotada neste projeto consiste na criação da função principal `ConverterTemperatura()`. Essa função recebe informações cruciais, como a unidade de medida original, a unidade desejada para a conversão e a temperatura a ser convertida. A estrutura de controle `switch` é utilizada para determinar a fórmula de conversão mais adequada.

## Testes Unitários

Os testes unitários são uma parte integral deste projeto, implementados utilizando o pacote `testing` do Golang. Esses testes são projetados para cobrir uma variedade de cenários, garantindo que o conversor de temperatura mantenha sua precisão e confiabilidade em diferentes condições.

## Trechos de Código

A seguir, destacamos alguns trechos de código representativos do projeto:

```go
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

// converterDeCelsius converte uma temperatura de Celsius para outra unidade de medida.
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
```

## Exemplo de Uso

```bash
Digite o seu nome: John
Olá, John!
Bem-vindo ao Conversor de Temperatura!

Escolha a unidade de medida da temperatura que você irá inserir:
[1] Celsius
[2] Fahrenheit
[3] Kelvin
[4] Cancelar e sair

Digite a temperatura em celsius: 100

Agora, escolha a unidade para a qual você deseja converter:
[1] Celsius
[2] Fahrenheit
[3] Kelvin
[4] Voltar para o menu anterior

100.00 celsius é igual a 212.00 fahrenheit

Deseja realizar outra conversão?
[1] Sim
[2] Encerrar programa
```

## Instalação

Certifique-se de ter o Go instalado. Em seguida, execute os seguintes comandos:

```bash
go get -u github.com/marcuscarvalhodev/conversorTemperatura
cd $GOPATH/src/github.com/marcuscarvalhodev/conversorTemperatura
go run main.go
```

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests.


## Aprendizado

Durante o desenvolvimento deste projeto, adquiri conhecimentos valiosos em:

- Entrada e saída em Golang.
- Estruturas de dados em Golang.
- Funções em Golang.
- Práticas de escrita de testes unitários em Golang.

## Conclusão

Esta jornada de desenvolvimento proporcionou uma experiência enriquecedora, consolidando meus conhecimentos e aprimorando minhas habilidades de programação. A importância da escrita de código testável foi evidenciada, garantindo a confiabilidade do conversor de temperatura em diversas situações.
