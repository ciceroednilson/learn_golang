package enderecos

import "strings"

// TipoDeEndereco retorna o tipo de endereço
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"Rua", "Avenida", "Estrada", "Rodovia"}

	primeiraPalavraDoEndereco := strings.Split(strings.ToLower(endereco), " ")[0]

	enderecoTemTipoValido := false

	for _, tipo := range tiposValidos {
		if strings.ToLower(tipo) == primeiraPalavraDoEndereco {
			enderecoTemTipoValido = true
		}
	}

	if enderecoTemTipoValido {
		return primeiraPalavraDoEndereco
	}

	return "Tipo inválido"
}
