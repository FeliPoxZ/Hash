// Pacote principal
package main

// Importação do pacote fmt para formatação de entrada/saída
import (
	"fmt"
)

// Estrutura Hash que contém um slice de VetorHash, um contador de quantidade, e um vetor para salvar índices adicionados
type Hash struct {
	Indices     []VetorHash // Slice de VetorHash
	Referencias []int       // Vetor para salvar índices adicionados
	Quantidade  int         // Contador de quantidade
}

// Estrutura VetorHash que contém um ponteiro para Dados e um verificador de colisão
type VetorHash struct {
	Dados_Usuario       *Dados // Ponteiro para Dados
	Verificador_colisao bool   // Verificador de colisão
}

// Estrutura Dados que contém informações do usuário e um ponteiro para o próximo Dados
type Dados struct {
	Nome     string // Nome do usuário
	Endereco string // Endereço do usuário
	Telefone string // Telefone do usuário
	Next     *Dados // Ponteiro para o próximo Dados
}

// Função principal
func main() {

	// Criação de um novo Hash
	var hash *Hash
	hash = CriaHash()

	// Inserção de dados no Hash
	InserirDados(hash, "Ana", "Rua", "TelefoneAna1")
	InserirDados(hash, "Felipe", "Rua", "Telefone1")
	InserirDados(hash, "Felipe", "Rua", "Telefone2")
	InserirDados(hash, "Quelli", "Rua", "TelefoneQuelli1")
	InserirDados(hash, "Ana", "Rua", "TelefoneAna2")
	InserirDados(hash, "Ana", "Rua", "TelefoneAna3")
	InserirDados(hash, "Felipe", "Rua", "Telefone3")
	InserirDados(hash, "Ana", "Rua", "TelefoneAna4")
	InserirDados(hash, "Ana", "Rua", "TelefoneAna5")
	InserirDados(hash, "Diego", "Rua", "Telefone")
	InserirDados(hash, "Eduardo", "Rua", "Telefone")
	InserirDados(hash, "Quelli", "Rua", "TelefoneQuelli2")
	InserirDados(hash, "Quelli", "Rua", "TelefoneQuelli3")

	// Mostra o Hash
	MostraHash(hash, hash.Referencias)

	// Deleta um elemento do Hash
	DeleteAllHash(hash, "Felipe")
	DeleteAllHash(hash, "Quelli")
	

	// Mostra o Hash após a deleção
	MostraHash(hash, hash.Referencias)
}

// Função para criar um novo Hash
func CriaHash() *Hash {
	// Cria um novo Hash com um slice de VetorHash de tamanho 10, um vetor de referências vazio e quantidade 0
	hash_Table := &Hash{Indices: make([]VetorHash, 10), Referencias: make([]int, 0), Quantidade: 0}
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
		Rehash(hash_table, Nome_input)
		Indice = Peso_strings(Nome_input, hash_table)
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
			current = current.Next
		}
		current.Next = Informacoes
	}

	// Incrementa a quantidade de dados no Hash
	hash_table.Quantidade++
	fmt.Println("\n", Nome_input)
	fmt.Println(hash_table.Referencias)
}

// Função para buscar dados no Hash
func BuscaHash(hash_table *Hash, Nome_search string) {

	// Calcula o índice onde os dados devem estar
	Indice := Peso_strings(Nome_search, hash_table)

	if len(hash_table.Indices) <= Indice {
		fmt.Println("\nNenhum dado encontrado")
		return
	}

	// Começa a busca no primeiro Dados no índice
	current := hash_table.Indices[Indice].Dados_Usuario

	// Se não há dados no índice, imprime uma mensagem
	if current == nil {
		fmt.Println("\nNenhum dado encontrado")
		return
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

// Função para reorganizar o Hash quando necessário
func Rehash(hash_table *Hash, novoNome string) {

	fmt.Print("\nIniciando rehash", "\n\n")
	Referencia := hash_table.Referencias
	hash_table.Referencias = make([]int, 0)
	hash_table.Quantidade = 0

	i := 0
	max := 100
	for FlagNovoPeso(hash_table, Referencia, novoNome) && i < max {
		i++
	}

	if i < max {
		// Cria um slice temporário para armazenar os dados
		tempDados := make([]*Dados, 0)
		for _, indice := range Referencia {
			Hash_Auxiliar := hash_table.Indices[indice]
			current := Hash_Auxiliar.Dados_Usuario
			for current != nil {
				tempDados = append(tempDados, current)
				current = current.Next
			}
			Hash_Auxiliar.Dados_Usuario = nil
			Hash_Auxiliar.Verificador_colisao = false
		}

		// Limpa a tabela hash
		hash_table.Indices = make([]VetorHash, len(hash_table.Indices))

		// Insere os dados novamente na tabela hash
		for _, dados := range tempDados {
			InserirDados(hash_table, dados.Nome, dados.Endereco, dados.Telefone)
		}

	} else {
		fmt.Println("Erro de rehash!")
		return
	}
	fmt.Println("Fim da rerash")
}

// Função que verifica se é necessário aumentar o vetor novamente
func FlagNovoPeso(hash_table *Hash, Referencia []int, novoNome string) bool {

	if Peso_strings(novoNome, hash_table) >= len(hash_table.Indices) {
		temporary := make([]VetorHash, Peso_strings(novoNome, hash_table)+1)
		copy(temporary, hash_table.Indices)
		hash_table.Indices = temporary
		// A função retorna true indicando necessidade de aumentar vetor novamente.
		return true
	}

	// O loop for percorre cada índice na lista de referências da tabela hash.
	for _, indice := range Referencia {
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

func DeleteHash(hash_table *Hash, Nome_Delete string, Telefone_Delete string) {
	Position := Peso_strings(Nome_Delete, hash_table)
	Hash := &hash_table.Indices[Position]

	count := 0

	current := Hash.Dados_Usuario
	var prev *Dados

	if Hash.Verificador_colisao {
		for current != nil {
			if current != nil && count == 0 && current.Nome == Nome_Delete && current.Telefone == Telefone_Delete {
				Hash.Dados_Usuario = current.Next
				count = 1
				return
			}else{
				if current == nil {
					return
				}
				if current.Nome == Nome_Delete && current.Telefone == Telefone_Delete {
					prev.Next = current.Next
					current = current.Next
				}
			}

			if current == nil{
				return
			}
			count = 1
			prev = current
			current = current.Next

		}

		Hash.Verificador_colisao = false
		current = Hash.Dados_Usuario
		for current != nil {
			if current.Next != nil && current.Nome != current.Next.Nome {
				Hash.Verificador_colisao = true
			}
			current = current.Next
		}
	} else {
		if current.Next == nil{
			Hash.Dados_Usuario = nil
			Hash.Verificador_colisao = false
			Referencias_auxiliar := make([]int, 0)

			for _, conteudo := range hash_table.Referencias{
				if conteudo != Position{
					Referencias_auxiliar = append(Referencias_auxiliar, conteudo)
				}
			}
			hash_table.Referencias = Referencias_auxiliar
			return
		}
		current = Hash.Dados_Usuario
		for current != nil {
			if current != nil && count == 0 && current.Nome == Nome_Delete && current.Telefone == Telefone_Delete {
				Hash.Dados_Usuario = current.Next
				count = 1
				return
			}else{
				if current == nil {
					return
				}
				if current.Nome == Nome_Delete && current.Telefone == Telefone_Delete {
					prev.Next = current.Next
					current = current.Next
				}
			}

			if current == nil{
				return
			}
			count = 1
			prev = current
			current = current.Next

		}
		

	}
}

// DeleteAllHash é uma função que remove todas as ocorrências de um nome específico (Nome_Delete) da tabela hash.
// A função começa calculando a posição do nome na tabela hash usando a função Peso_strings.
func DeleteAllHash(hash_table *Hash, Nome_Delete string) {

	// Calcula a posição do nome na tabela hash.
	Position := Peso_strings(Nome_Delete, hash_table)
	// Cria um ponteiro para o VetorHash na posição calculada.
	Hash := &hash_table.Indices[Position]

	// Verifica se há colisão na posição calculada.
	if Hash.Verificador_colisao {
		// Se houver colisão, a função percorre a lista ligada na posição.
		current := Hash.Dados_Usuario
		var prev *Dados

		// Percorre a lista ligada.
		for current != nil {
			// Se o nome do usuário atual for igual ao nome a ser deletado.
			if current.Nome == Nome_Delete {
				// Se prev não for nil, atualiza o próximo de prev para ser o próximo de current.
				// Caso contrário, atualiza o Dados_Usuario de Hash para ser o próximo de current.
				if prev != nil {
					prev.Next = current.Next
				} else {
					Hash.Dados_Usuario = current.Next
				}
			} else {
				// Se o nome do usuário atual não for igual ao nome a ser deletado, atualiza prev para ser current.
				prev = current
			}
			// Se current não for nil, atualiza current para ser o próximo de current.
			if current != nil {
				current = current.Next
			}
		}

		// Atualiza current para ser Dados_Usuario de Hash.
		current = Hash.Dados_Usuario
		// Define o verificador de colisão na posição como falso.
		hash_table.Indices[Position].Verificador_colisao = false
		// Percorre a lista ligada novamente.
		for current != nil {
			// Se o próximo de current não for nil e o nome de current não for igual ao nome do próximo de current.
			// Define o verificador de colisão na posição como verdadeiro.
			if current.Next != nil && current.Nome != current.Next.Nome {
				hash_table.Indices[Position].Verificador_colisao = true
			}
			// Atualiza current para ser o próximo de current.
			current = current.Next
		}

	} else {
		// Se não houver colisão na posição, define Dados_Usuario de Hash e o verificador de colisão como nil e falso, respectivamente.
		Hash.Dados_Usuario = nil
		Hash.Verificador_colisao = false
		// Cria um slice auxiliar de referências com o mesmo tamanho do slice de referências da tabela hash.
		Referencias_auxiliar := make([]int, len(hash_table.Referencias))
		// Percorre o slice de referências da tabela hash.
		for _, Conteudo := range hash_table.Referencias {
			// Se o conteúdo não for igual à posição, adiciona o conteúdo ao slice auxiliar de referências.
			if Conteudo != Position {
				Referencias_auxiliar = append(Referencias_auxiliar, Conteudo)
			}
		}
		// Atualiza o slice de referências da tabela hash para ser o slice auxiliar de referências.
		hash_table.Referencias = Referencias_auxiliar
	}
}

// Função para mostrar o Hash
func MostraHash(hash_table *Hash, Referencias []int) {

	
	// Para cada índice na lista de referências
	for _, indices := range Referencias {
		Hash := &hash_table.Indices[indices]
		if Hash.Dados_Usuario != nil {
			if Hash.Dados_Usuario.Next == nil {
				fmt.Printf("\nNome: %s", Hash.Dados_Usuario.Nome)
			} else {
				fmt.Printf("\nNome: %s -> ", Hash.Dados_Usuario.Nome)
			}
			current := Hash.Dados_Usuario.Next
			// Percorre a lista ligada no índice, imprimindo os nomes
			for current != nil {
				if current.Next == nil {
					fmt.Printf(" %s", current.Nome)
				} else {
					fmt.Printf("%s -> ", current.Nome)
				}
				current = current.Next
			}
		}
	}

	fmt.Println()
}
