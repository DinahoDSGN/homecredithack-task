package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllShops(c *gin.Context) {
	data, _ := h.services.Shop.GetAll()

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
