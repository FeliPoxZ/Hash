package main

import (
	"fmt"
)

// Estrutura Hash que contém um slice de VetorHash, um contador de quantidade, e um vetor para salvar indices adicionados
type Hash struct {
	Indices     []VetorHash
	Referencias []int
	Quantidade  int
}

// Estrutura VetorHash que contém um ponteiro para Dados e um verificador de colisão
type VetorHash struct {
	Dados_Usuario       *Dados
	Verificador_colisao bool
}

// Estrutura Dados que contém informações do usuário e um ponteiro para o próximo Dados
type Dados struct {
	Nome     string
	Endereco string
	Telefone string
	Next     *Dados
}

func main() {

	var hash *Hash

	// Cria um novo Hash
	hash = CriaHash()

	// Insere dados no Hash
	InserirDados(hash, "Felipe", "Rua", "Telefone")
	InserirDados(hash, "Felipe", "Rua", "Telefone")
	InserirDados(hash, "Ana", "Rua", "Telefone")

	// Busca dados no Hash
	BuscaHash(hash, "Felipe")

}

// Função para criar um novo Hash
func CriaHash() *Hash {

	// Cria um novo Hash com um slice de VetorHash de tamanho 250
	hash_Table := &Hash{Indices: make([]VetorHash, 250), Referencias: make([]int, 0), Quantidade: 0}

	return hash_Table
}

// Função para calcular o peso de uma string
func Peso_strings(nome string, hash_table *Hash) int {

	var Peso int
	Grau := len(nome)

	// Calcula o peso da string
	Somatoria := 0
	for _, Letra := range nome {
		Somatoria += int(Letra) * Grau
		Grau--
	}

	Peso = Somatoria

	// Calcula o índice baseado no peso e no tamanho do slice
	Resto := Peso % (len(hash_table.Indices) + 1)
	println(Peso)
	return Resto
}

// Função para inserir dados no Hash
func InserirDados(hash_table *Hash, Nome_input string, Endereco_input string, Telefone_input string) {

	// Calcula o índice onde os dados devem ser inseridos
	Indice := Peso_strings(Nome_input, hash_table)

	// Cria um novo Dados com as informações do usuário
	Informacoes := &Dados{Nome: Nome_input, Endereco: Endereco_input, Telefone: Telefone_input}

	// Se o índice for maior do que o tamanho do slice, redimensiona o slice
	if len(hash_table.Indices) <= Indice {
		temporary := make([]VetorHash, len(hash_table.Indices)*2)
		copy(temporary, hash_table.Indices)
		hash_table.Indices = temporary
	}

	// Cria um alias para o VetorHash no índice
	Hash := &hash_table.Indices[Indice]

	// Se o VetorHash no índice não contém dados, insere os dados
	// Se já contém dados, insere os novos dados no final da lista ligada
	if Hash.Dados_Usuario == nil {
		Hash.Dados_Usuario = Informacoes
		hash_table.Referencias = append(hash_table.Referencias, Indice)
	} else {
		current := Hash.Dados_Usuario
		for current.Next != nil {
			if Informacoes.Nome != current.Nome {
				Hash.Verificador_colisao = true
			}
			current = Hash.Dados_Usuario.Next
		}
		current.Next = Informacoes
	}

	// Incrementa a quantidade de dados no Hash
	hash_table.Quantidade++
	fmt.Println(hash_table.Referencias)
}

// Função para buscar dados no Hash
func BuscaHash(hash_table *Hash, Nome_search string) {

	// Calcula o índice onde os dados devem estar
	Indice := Peso_strings(Nome_search, hash_table)

	// Começa a busca no primeiro Dados no índice
	current := hash_table.Indices[Indice].Dados_Usuario

	// Se não há dados no índice, imprime uma mensagem
	if current == nil {
		fmt.Println("Nenhum dado encontrado")
	}

	// Percorre a lista ligada no índice, imprimindo os dados que correspondem ao nome buscado
	i := 0
	for current != nil {
		if current.Nome == Nome_search {
			i++
			fmt.Printf("\nNome: %s (%d)", current.Nome, i)
			fmt.Printf("\nEndereço: %s", current.Endereco)
			fmt.Printf("\nTelefone: %s\n", current.Telefone)
		}
		current = current.Next
	}
}

func Rehash(hash_table *Hash) {

	// Novo
	i := 0
	max := 100
	for FlagNovoPeso(hash_table) && i < max {
		i++
	}

	// If novo
	if i < max {
		for _, indice := range hash_table.Referencias {
			Hash := hash_table.Indices[indice]
			if Hash.Verificador_colisao {

			} else {
				NovoIndice := Peso_strings(Hash.Dados_Usuario.Nome, hash_table)
				hash_table.Indices[NovoIndice].Dados_Usuario = Hash.Dados_Usuario
				hash_table.Indices[NovoIndice].Verificador_colisao = false

				//Novo
				Hash.Dados_Usuario = nil
			}
		}
	} else {
		fmt.Println("Erro de rehash!")
		return
	}

}

// Novo
// FlagNovoPeso é uma função que recebe uma tabela hash e retorna um booleano.
func FlagNovoPeso(hash_table *Hash) bool {

	// O loop for percorre cada índice na lista de referências da tabela hash.
	for _, indice := range hash_table.Referencias {
		// Hash é um alias que armazena o valor no índice atual da tabela hash.
		Hash := hash_table.Indices[indice]
		// Se o verificador de colisão do Hash for verdadeiro, o código dentro deste bloco if será executado.
		if Hash.Verificador_colisao {
			current := Hash.Dados_Usuario
			// Este loop for continuará enquanto current não for nil.
			for current != nil {
				NovoIndice := Peso_strings(current.Nome, hash_table)
				// Se o tamanho da lista de índices da tabela hash for menor ou igual a NovoIndice, ajusta o tamanho do vetor.
				if len(hash_table.Indices) <= NovoIndice {
					temporary := make([]VetorHash, NovoIndice+1)
					copy(temporary, hash_table.Indices)
					hash_table.Indices = temporary
					// A função retorna true indicando necessidade de aumentar vetor novamente.
					return true
				}
				// current é atualizado para ser o próximo valor na lista ligada.
				current = current.Next
			}
		} else {
			NovoIndice := Peso_strings(Hash.Dados_Usuario.Nome, hash_table)
			// Se o tamanho da lista de índices da tabela hash for menor ou igual a NovoIndice, o código dentro deste bloco if será executado.
			if len(hash_table.Indices) <= NovoIndice {
				temporary := make([]VetorHash, NovoIndice+1)
				copy(temporary, hash_table.Indices)
				hash_table.Indices = temporary
				// A função retorna true indicando necessidade de aumentar vetor novamente.
				return true
			}
		}
	}
	// Se o loop for terminar sem retornar true, a função retornará false, indicando que não é preciso aumentar vetor novamente.
	return false
}
