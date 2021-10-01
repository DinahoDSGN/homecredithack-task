package repository

import (
	"HomeCreditHack/pkg/models"
	"gorm.io/gorm"
)

type CreditBankRepository struct {
	db *gorm.DB
}

func NewCreditBankRepository(db *gorm.DB) *CreditBankRepository {
	return &CreditBankRepository{db: db}
}

func (r *CreditBankRepository) Create(data models.CreditCard) (models.CreditCard, error) {
	return models.CreditCard{}, nil
}

func (r *CreditBankRepository) GetAll() ([]models.CreditCard, error) {
	var records []models.CreditCard

	r.db.Model(&models.CreditCard{}).Find(&records)

	return records, nil
}

func (r *CreditBankRepository) GetList(id int) (models.CreditCard, error) {
	return models.CreditCard{}, nil
}

func (r *CreditBankRepository) Delete(id int) (models.CreditCard, error) {
	return models.CreditCard{}, nil
}

func (r *CreditBankRepository) Update(id int, data models.CreditCard) (models.CreditCard, error) {
	return models.CreditCard{}, nil
}
