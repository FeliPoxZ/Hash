package main

import (
	"fmt"
)

type Hash struct {
	Indices []*Dados
	Tamanho int
}

type Dados struct {
	Nome     string
	Endereco string
	Telefone string
	Next     *Dados
	Prev     *Dados
}

func main() {

	var hash *Hash

	hash = CriaHash()

	InserirDados(hash)
	fmt.Println(hash.Indices[7])

}

func CriaHash() *Hash {

	hash_Table := &Hash{Indices: nil, Tamanho: 0}

	hash_Table.Indices = make([]*Dados, 0) //Indices aponta agora para um slice
	return hash_Table
}

func Peso_strings(nome string) int {

	var Peso int
	Grau := len(nome)

	Somatoria := 0
	for _, Letra := range nome {
		Somatoria += int(Letra) * Grau
		Grau--
	}

	Peso = Somatoria
	return Peso % 10
}

func InserirDados(hash_table *Hash) {

	var Nome_input, Endereco_input, Telefone_Input string

	fmt.Println("Insira seu nome, seu enderço e seu telefone")
	fmt.Scanf("%s %s %s", &Nome_input, &Endereco_input, &Telefone_Input)

	Indice_peso := Peso_strings(Nome_input)

	if len(hash_table.Indices) < Indice_peso {

		temporary := make([]*Dados, Indice_peso+1)
		copy(temporary, hash_table.Indices)
		hash_table.Indices = temporary
		hash_table.Tamanho = len(hash_table.Indices)
	}

	Informacoes := &Dados{Nome: Nome_input, Endereco: Endereco_input, Telefone: Telefone_Input}

	hash_table.Indices[Indice_peso] = Informacoes
}
