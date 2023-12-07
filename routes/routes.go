package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sajjad3k/controllers"
)

func SetRoutes(router *gin.Engine) {
	// Produtos
	router.POST("/produtos/create", controllers.AdicionarProdutoController)
	router.GET("/produtos", controllers.ListarTodosProdutosController)
	router.GET("/produtos/edit/get/:manutID", controllers.BuscarProdutoController)
	router.PUT("/produtos/edit/update/:manutID", controllers.AtualizarProdutoController)
	router.DELETE("/produtos/delete/:manutID", controllers.DeletarProdutoController)
}
