package main

import (
	"go_abacatepay/handlers"
	"go_abacatepay/internal/abacatepayclient"

	"github.com/gin-gonic/gin"
)

func main() {
	// Cria o cliente AbacatePay com URL base e token
	client := abacatepayclient.NewClient(
		"https://api.abacatepay.com",
		"",
	)

	// Cria o router gin
	r := gin.Default()

	// Registra as rotas e passa o client para os handlers
	r.POST("/clientes", handlers.CriarClienteHandler(client))
	r.POST("/pagamentos", handlers.CriarPagamentoHandler(client))

	// Roda o servidor na porta 8000
	r.Run(":8000")
}
