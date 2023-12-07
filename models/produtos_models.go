package models

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Produto - Estrutura para representar um produto
type Produto struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	Nome      string             `bson:"nome" json:"nome"`
	Valor     string             `bson:"valor" json:"valor"`
	Descricao string             `bson:"descricao" json:"descricao"`
	Categoria string             `bson:"categoria" json:"categoria"`
	CreatedAt time.Time          `bson:"created_at" json:"createdAt"`
}

var produtos []Produto

// init - Função inicializadora para criar a lista de produtos
func init() {
	produtos = make([]Produto, 0)
}

// AdicionarProduto adiciona um novo produto ao repositório
func AdicionarProduto(nome, valor, descricao, categoria string) Produto {
	produto := Produto{
		ID:        primitive.NewObjectID(),
		Nome:      nome,
		Valor:     valor,
		Descricao: descricao,
		Categoria: categoria,
		CreatedAt: time.Now(),
	}

	produtos = append(produtos, produto)
	return produto
}

// BuscarProduto retorna um produto pelo ID
func BuscarProduto(id string) (Produto, error) {
	for _, produto := range produtos {
		if produto.ID.Hex() == id {
			return produto, nil
		}
	}
	return Produto{}, errors.New("produto não encontrado")
}

// AtualizarProduto atualiza um produto existente
func AtualizarProduto(id string, nome, valor, descricao, categoria string) error {
	for i, produto := range produtos {
		if produto.ID.Hex() == id {
			produtos[i].Nome = nome
			produtos[i].Valor = valor
			produtos[i].Descricao = descricao
			produtos[i].Categoria = categoria
			return nil
		}
	}
	return errors.New("produto não encontrado")
}

// DeletarProduto remove um produto pelo ID
func DeletarProduto(id string) error {
	for i, produto := range produtos {
		if produto.ID.Hex() == id {
			produtos = append(produtos[:i], produtos[i+1:]...)
			return nil
		}
	}
	return errors.New("produto não encontrado")
}

// ListarTodosProdutos retorna todos os produtos
func ListarTodosProdutos() []Produto {
	return produtos
}
