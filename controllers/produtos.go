package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajjad3k/models"
	"log"
	"net/http"
)

func AdicionarProdutoController(c *gin.Context) {
	var novoProduto models.Produto
	if err := c.ShouldBindJSON(&novoProduto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	produtoAdicionado := models.AdicionarProduto(novoProduto.Nome, novoProduto.Valor, novoProduto.Descricao, novoProduto.Categoria)
	c.JSON(http.StatusCreated, produtoAdicionado) // Código de status para criação bem-sucedida
}

func AtualizarProdutoController(c *gin.Context) {
	var produtoAtualizado models.Produto
	if err := c.ShouldBindJSON(&produtoAtualizado); err != nil {
		log.Printf("Erro ao decodificar o corpo da requisição: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.AtualizarProduto(produtoAtualizado.ID.Hex(), produtoAtualizado.Nome, produtoAtualizado.Valor, produtoAtualizado.Descricao, produtoAtualizado.Categoria)
	if err != nil {
		log.Printf("Erro ao atualizar o produto: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar o produto"})
		return
	}

	c.JSON(http.StatusOK, produtoAtualizado)
}

func DeletarProdutoController(c *gin.Context) {
	var produto models.Produto
	if err := c.ShouldBindJSON(&produto); err != nil {
		log.Printf("Erro ao decodificar o corpo da requisição: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.DeletarProduto(produto.ID.Hex())
	if err != nil {
		log.Printf("Erro ao deletar o produto: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar o produto"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produto deletado com sucesso"})
}

func BuscarProdutoController(c *gin.Context) {
	var produto models.Produto
	if err := c.ShouldBindJSON(&produto); err != nil {
		log.Printf("Erro ao decodificar o corpo da requisição: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	produtoEncontrado, err := models.BuscarProduto(produto.ID.Hex())
	if err != nil {
		log.Printf("Erro ao buscar o produto: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}

	c.JSON(http.StatusOK, produtoEncontrado)
}

func ListarTodosProdutosController(c *gin.Context) {
	produtos := models.ListarTodosProdutos()
	c.JSON(http.StatusOK, produtos)
}
