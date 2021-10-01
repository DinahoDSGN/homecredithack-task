package repository

import (
	"HomeCreditHack/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ShopRepository struct {
	db *gorm.DB
}

func NewShopRepository(db *gorm.DB) *ShopRepository {
	return &ShopRepository{db: db}
}

func (database ShopRepository) Create(data models.Shop) (models.Shop, error) {
	return models.Shop{}, nil
}

func (database ShopRepository) GetAll() ([]models.Shop, error) {
	var records []models.Shop

	database.db.Preload(clause.Associations).Model(&models.Shop{}).Find(&records)

	return records, nil
}

func (database ShopRepository) GetList(id int) (models.Shop, error) {
	return models.Shop{}, nil
}

func (database ShopRepository) Delete(id int) (models.Shop, error) {
	return models.Shop{}, nil
}

func (database ShopRepository) Update(id int, data models.Shop) (models.Shop, error) {
	return models.Shop{}, nil
}
