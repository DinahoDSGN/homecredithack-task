package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllOrders(c *gin.Context) {
	data, _ := h.services.Order.GetAll()

	c.JSON(http.StatusOK, gin.H{
		"data" : data,
	})
}
