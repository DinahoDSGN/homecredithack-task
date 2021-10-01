package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Order_item struct {
	Id uint
	ProductId uint
	Product *Product `gorm:"foreignKey:Id"`
	OrderId uint
	Order *Order `gorm:"foreignKey:Id"`
}

type Order struct {
	Id uint
	ShopId uint
	Shop *Shop `gorm:"foreignKey:ShopId"`
	Status bool
	Amount uint
	Total float64
	Mobile string
}

func (order *Order_item) AfterCreate(tx *gorm.DB) (err error){
	var records Order_item
	tx.Preload(clause.Associations).Raw("SELECT * FROM order_items WHERE id = ?", order.Id).Find(&records)

	order.Order.Total = float64(uint(order.Product.Price) * order.Order.Amount)

	return nil
}