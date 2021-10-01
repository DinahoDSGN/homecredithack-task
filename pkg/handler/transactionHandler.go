package handler

import (
	"HomeCreditHack/pkg/database"
	"HomeCreditHack/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

func (h *Handler) GetAllTransactions(c *gin.Context) {
	var records []models.Order_item
	database.DB.Preload(clause.Associations).Model(&models.Order_item{}).Find(&records)
	c.JSON(http.StatusOK, gin.H{
		"data" : records,
	})
}
