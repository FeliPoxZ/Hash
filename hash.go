package main

import (
	"fmt"
)

//map := make(map[int]*Dados)

type Hash struct {
	Indices    []VetorHash
	Quantidade int
}

type VetorHash struct {
	Dados_Usuario     *Dados
	Verificador_colisao bool
}

type Dados struct {
	Nome     string
	Endereco string
	Telefone string
	Next     *Dados
	//Prev     *Dados
}

func main() {

	var hash *Hash

	hash = CriaHash()

	InserirDados(hash, "Felipe", "Rua", "Telefone")
	InserirDados(hash, "Felipe", "Rua", "Telefone")
	InserirDados(hash, "Ana", "Rua", "Telefone")
	BuscaHash(hash, "Felipe")

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

	//Violeta

	Peso = Somatoria
	retorno := Peso % (len(hash_tb.Indices) + 1)
	fmt.Print("Nome: ", nome, "\n", "Peso: ", Peso,"\n" ,"Resto: ", retorno, "\n\n")
	return retorno
}

func InserirDados(hash_table *Hash, Nome_input string, Endereco_input string, Telefone_input string) {

	Indice := Peso_strings(Nome_input, hash_table)
	
	Informacoes := &Dados{Nome: Nome_input, Endereco: Endereco_input, Telefone: Telefone_input}
	//Indice := Indice
	
	//fmt.Println("Indice:", Indice)
	//fmt.Println("Length of Indices:", len(hash_table.Indices))
	
	if len(hash_table.Indices) <= Indice { 
		//fmt.Println("Resizing Indices...")
		temporary := make([]VetorHash, len(hash_table.Indices)*2)
		copy(temporary, hash_table.Indices)
		hash_table.Indices = temporary
	}
	
	Hash := &hash_table.Indices[Indice] //alias

	if Hash.Dados_Usuario == nil {
		Hash.Dados_Usuario = Informacoes
	} else {
		current := Hash.Dados_Usuario
		for current.Next != nil {
			if(Informacoes.Nome != current.Nome){
				Hash.Verificador_colisao = true
			}
			current = Hash.Dados_Usuario.Next
		}
		current.Next = Informacoes
	}

	hash_table.Quantidade++
}

func BuscaHash(Hash_Table *Hash, Nome_search string) {

	Indice := Peso_strings(Nome_search, Hash_Table)

	current := Hash_Table.Indices[Indice].Dados_Usuario
	i := 0

	if current == nil {
		fmt.Println("Nenhum dado encontrado")
	}

	for current != nil {
		if current.Nome == Nome_search {
			i ++
			fmt.Printf("\nNome: %s (%d)", current.Nome, i)
			fmt.Printf("\nEndereço: %s", current.Endereco)
			fmt.Printf("\nTelefone: %s\n", current.Telefone)
		}
		current = current.Next
	}

}

