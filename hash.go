package main

import (
	"fmt"
)

// Estrutura Hash que contém um slice de VetorHash, um contador de quantidade, e um vetor para salvar indices adicionados
type Hash struct {
	Indices    []VetorHash
	Referencias []int
	Quantidade int
}

// Estrutura VetorHash que contém um ponteiro para Dados e um verificador de colisão
type VetorHash struct {
	Dados_Usuario     *Dados
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
func Peso_strings(nome string, hash_tb *Hash) int {

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
	Resto := Peso % (len(hash_tb.Indices) + 1)

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
			if(Informacoes.Nome != current.Nome){
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
func BuscaHash(Hash_Table *Hash, Nome_search string) {

	// Calcula o índice onde os dados devem estar
	Indice := Peso_strings(Nome_search, Hash_Table)

	// Começa a busca no primeiro Dados no índice
	current := Hash_Table.Indices[Indice].Dados_Usuario

	// Se não há dados no índice, imprime uma mensagem
	if current == nil {
		fmt.Println("Nenhum dado encontrado")
	}

	// Percorre a lista ligada no índice, imprimindo os dados que correspondem ao nome buscado
	i := 0
	for current != nil {
		if current.Nome == Nome_search {
			i++
			fmt.Printf("\nNome: %s (%d)", current.Nome,i)
			fmt.Printf("\nEndereço: %s", current.Endereco)
			fmt.Printf("\nTelefone: %s\n", current.Telefone)
		}
		current = current.Next
	}
}
