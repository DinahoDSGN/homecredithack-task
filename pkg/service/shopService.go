package service

import (
	"HomeCreditHack/pkg/models"
	"HomeCreditHack/pkg/repository"
)

type ShopService struct {
	repo repository.Shop
}

func NewShopService(repo repository.Shop) *ShopService {
	return &ShopService{repo: repo}
}

func (s ShopService) Create(data models.Shop) (models.Shop, error) {
	return models.Shop{}, nil
}

func (s ShopService) GetAll() ([]models.Shop, error) {
	records, _ := s.repo.GetAll()

	return records, nil
}

func (s ShopService) GetList(id int) (models.Shop, error) {
	return models.Shop{}, nil
}

func (s ShopService) Delete(id int) (models.Shop, error) {
	return models.Shop{}, nil
}

func (s ShopService) Update(id int, data models.Shop) (models.Shop, error) {
	return models.Shop{}, nil
}
