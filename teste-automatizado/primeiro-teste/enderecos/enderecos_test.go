// teste de unidade
package enderecos

import (
	"strings"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoEndereco(t *testing.T) {

	t.Parallel()
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua Azul", "Rua"},
		{"Avenida Azul", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		/*	{"Praça das rosas", "Praça"}, */
		{"Estrada das Pedras", "Estrada"},
	}
	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != strings.ToLower(cenario.retornoEsperado) {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", tipoDeEnderecoRecebido, cenario.retornoEsperado)
		}
	}
}

func TestSuccess(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Error("Error")
	}
}
