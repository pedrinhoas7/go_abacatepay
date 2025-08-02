package handlers

import (
	"net/http"

	"go_abacatepay/internal/abacatepayclient"

	"github.com/gin-gonic/gin"
)

func CriarClienteHandler(client *abacatepayclient.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req abacatepayclient.CustomerRequest
		req.Cellphone = c.Query("cellphone")
		req.Email = c.Query("email")
		req.Name = c.Query("name")
		req.TaxId = c.Query("taxId")
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"erro1": err.Error()})
			return
		}
		// Aqui o req deve estar preenchido corretamente
		resp, err := client.CreateCustomer(req)
		if err != nil {
			c.JSON(500, gin.H{"erro2": err.Error()})
			return
		}
		c.JSON(200, resp)
	}
}

func CriarPagamentoHandler(client *abacatepayclient.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req abacatepayclient.BillingRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
			return
		}

		resp, err := client.CreateBilling(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}
