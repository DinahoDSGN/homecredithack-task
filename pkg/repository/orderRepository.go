package repository

import (
	"HomeCreditHack/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderRepository struct {
	db *gorm.DB
}

func (database *OrderRepository) Create(data models.Order) (models.Order, error) {
	return models.Order{}, nil
}

func (database *OrderRepository) GetAll() ([]models.Order, error) {
	var records []models.Order

	database.db.Preload(clause.Associations).Model(&models.Order{}).Find(&records)

	return records, nil
}

func (database *OrderRepository) GetList(id int) (models.Order, error) {
	return models.Order{}, nil
}

func (database *OrderRepository) Delete(id int) (models.Order, error) {
	return models.Order{}, nil
}

func (database *OrderRepository) Update(id int, data models.Order) (models.Order, error) {
	return models.Order{}, nil
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}
