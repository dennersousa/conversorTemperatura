package conversor

import (
	"bytes"
	"testing"
)

func TestConverterTemperatura(t *testing.T) {
	tests := []struct {
		unidadeOriginal     string
		unidadeConvertida   string
		temperaturaInserida float64
		expectedOutput      string
	}{
		{"celsius", "fahrenheit", 0, "0.00 celsius é igual a 32.00 fahrenheit\n"},
		{"fahrenheit", "celsius", 32, "32.00 fahrenheit é igual a 0.00 celsius\n"},
		{"kelvin", "celsius", 273.15, "273.15 kelvin é igual a 0.00 celsius\n"},
	}

	for _, test := range tests {
		var buf bytes.Buffer
		ConverterTemperatura(test.unidadeOriginal, test.unidadeConvertida, test.temperaturaInserida, &buf)
		output := buf.String()
		if output != test.expectedOutput {
			t.Errorf("Conversão incorreta. Esperado:\n%s\nObtido:\n%s", test.expectedOutput, output)
		}
	}
}
