package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllCreditCards(c *gin.Context) {
	data, _ := h.services.CreditCard.GetAll()

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
