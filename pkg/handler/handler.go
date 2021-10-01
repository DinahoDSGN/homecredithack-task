package handler

import (
	"HomeCreditHack/pkg/service"
	"github.com/gin-gonic/gin"
	"log"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(route *gin.Engine) *gin.Engine {
	router := route

	api := router.Group("/api")
	{
		creditcard := api.Group("/cc")
		{
			creditcard.POST("/create")
			creditcard.GET("/", h.GetAllCreditCards)
			creditcard.GET("/:id")
			creditcard.DELETE("/:id")
			creditcard.PUT("/:id")
		}
		shop := api.Group("/shop")
		{
			shop.POST("/create")
			shop.GET("/", h.GetAllShops)
			shop.GET("/:id")
			shop.DELETE("/:id")
			shop.PUT("/:id")
		}

		transaction := api.Group("/transaction")
		{
			transaction.GET("/", h.GetAllTransactions)
		}

		order := api.Group("/order")
		{
			order.GET("/", h.GetAllOrders)
		}
	}

	log.Println("Routes initialized")

	return router
}


