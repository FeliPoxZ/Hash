package main

import (
	"fmt"
)

//map := make(map[int]*Dados)

type Hash struct {
	Indices []VetorHash
	Quantidade int
}

type VetorHash struct {
	Dados_Usuario *Dados
	Nome string
	Referencia_Indice int
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

	InserirDados(hash, "Felipe", "Rua", "Telefone")
	fmt.Println(hash.Indices[7])

}

func CriaHash() *Hash {

	hash_Table := &Hash{Indices: nil, Quantidade: 0}

	hash_Table.Indices = make([]VetorHash, 10) //Indices aponta agora para um slice
	return hash_Table
}

func Peso_strings(nome string, hash_tb *Hash) int {

	var Peso int
	Grau := len(nome)

	Somatoria := 0
	for _, Letra := range nome {
		Somatoria += int(Letra) * Grau
		Grau--
	}

	Peso = Somatoria
	return Peso % (len(hash_tb.Indices) + 1)
}

func InserirDados(hash_table *Hash, Nome_input string, Endereco_input string, Telefone_input string) {

	Indice_peso := Peso_strings(Nome_input, hash_table)
	fmt.Println(Indice_peso)

 	if len(hash_table.Indices) < Indice_peso {

		temporary := make([]VetorHash, Indice_peso+1)
		copy(temporary, hash_table.Indices)
		hash_table.Indices = temporary
	} 

	Informacoes := &Dados{Nome: Nome_input, Endereco: Endereco_input, Telefone: Telefone_input}


	hash_table.Indices[Indice_peso].Dados_Usuario = Informacoes
	hash_table.Indices[Indice_peso].Referencia_Indice = Indice_peso
	hash_table.Quantidade++
}
