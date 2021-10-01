package service

import (
	"HomeCreditHack/pkg/models"
	"HomeCreditHack/pkg/repository"
)

type CreditCardService struct {
	repo repository.CreditCard
}

func NewCreditCardService(repo repository.CreditCard) *CreditCardService {
	return &CreditCardService{repo: repo}
}

func (s *CreditCardService) Create(data models.CreditCard) (models.CreditCard, error) {
	return models.CreditCard{}, nil
}

func (s *CreditCardService) GetAll() ([]models.CreditCard, error) {
	records, _ := s.repo.GetAll()

	return records, nil
}

func (s *CreditCardService) GetList(id int) (models.CreditCard, error) {
	return models.CreditCard{}, nil
}

func (s *CreditCardService) Delete(id int) (models.CreditCard, error) {
	return models.CreditCard{}, nil
}

func (s *CreditCardService) Update(id int, data models.CreditCard) (models.CreditCard, error) {
	return models.CreditCard{}, nil
}
