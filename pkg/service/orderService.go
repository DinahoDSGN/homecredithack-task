package service

import (
	"HomeCreditHack/pkg/models"
	"HomeCreditHack/pkg/repository"
)

type OrderService struct {
	repos repository.Order
}

func (s *OrderService) Create(data models.Order) (models.Order, error) {
	return models.Order{}, nil
}

func (s *OrderService) GetAll() ([]models.Order, error) {
	records, _ := s.repos.GetAll()

	return records, nil
}

func (s *OrderService) GetList(id int) (models.Order, error) {
	return models.Order{}, nil
}

func (s *OrderService) Delete(id int) (models.Order, error) {
	return models.Order{}, nil
}

func (s *OrderService) Update(id int, data models.Order) (models.Order, error) {
	return models.Order{}, nil
}

func NewOrderService(repos repository.Order) *OrderService {
	return &OrderService{repos: repos}
}
