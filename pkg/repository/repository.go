package repository

import (
	"HomeCreditHack/pkg/models"
	"gorm.io/gorm"
)

type CreditCard interface {
	Create(data models.CreditCard) (models.CreditCard, error)
	GetAll() ([]models.CreditCard, error)
	GetList(id int) (models.CreditCard, error)
	Delete(id int) (models.CreditCard, error)
	Update(id int, data models.CreditCard) (models.CreditCard, error)
}

type Shop interface {
	Create(data models.Shop) (models.Shop, error)
	GetAll() ([]models.Shop, error)
	GetList(id int) (models.Shop, error)
	Delete(id int) (models.Shop, error)
	Update(id int, data models.Shop) (models.Shop, error)
}

type Order interface {
	Create(data models.Order) (models.Order, error)
	GetAll() ([]models.Order, error)
	GetList(id int) (models.Order, error)
	Delete(id int) (models.Order, error)
	Update(id int, data models.Order) (models.Order, error)
}

type Repository struct {
	CreditCard
	Shop
	Order
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		CreditCard: NewCreditBankRepository(db),
		Shop:       NewShopRepository(db),
		Order:      NewOrderRepository(db),
	}
}
