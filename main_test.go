package main

import (
	"bytes"
	"os/exec"
	"regexp"
	"strings"
	"testing"
)

func TestMainFunction(t *testing.T) {
	// Simulando a entrada do usuário durante o teste
	input := "John\n1\n100\n2\n2\n"
	expectedOutput := `Digite o seu nome: Olá, John!
Bem-vindo ao Conversor de Temperatura!
Escolha a unidade de medida da temperatura que você irá inserir:
[1] Celsius
[2] Fahrenheit
[3] Kelvin
[4] Cancelar e sair
Digite a temperatura em celsius: Agora, escolha a unidade para a qual você deseja converter:
[1] Celsius
[2] Fahrenheit
[3] Kelvin
[4] Voltar para o menu anterior
100.00 celsius é igual a 212.00 fahrenheit
Deseja realizar outra conversão?
[1] Sim
[2] Encerrar programa
Programa encerrado. Até logo, John!`

	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin = strings.NewReader(input)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		t.Fatalf("Erro ao executar o teste: %v\n%s", err, stderr.String())
	}

	output := strings.TrimSpace(stdout.String())
	expectedOutput = strings.TrimSpace(expectedOutput)

	// Remover espaços extras no nome do usuário
	re := regexp.MustCompile(`Olá,\s*John!`)
	expectedOutput = re.ReplaceAllString(expectedOutput, "Olá, John!")

	if output != expectedOutput {
		t.Errorf("A saída esperada não correspondeu à saída real.\nEsperado:\n%s\nRecebido:\n%s", expectedOutput, output)
	}
}
