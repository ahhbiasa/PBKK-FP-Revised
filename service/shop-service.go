package service

import (
	entities "PBKK-FP-Revised/entities"
	"PBKK-FP-Revised/repository"
)

type ShopService interface {
	Save(shop entities.Shop) error
	Update(shop entities.Shop)
	Delete(shop entities.Shop)
	FindAll() ([]entities.Shop, error) // Include error in the return value
	GetShopByID(id int) (entities.Shop, error)
}

type shopService struct {
	ShopRepository repository.ShopRepository
}

func NewShopService(repo repository.ShopRepository) ShopService {
	return &shopService{
		ShopRepository: repo,
	}
}

func (service *shopService) Save(shop entities.Shop) error {
	err := service.ShopRepository.Save(shop)
	if err != nil {
		return err
	}
	return nil
}

func (service *shopService) FindAll() ([]entities.Shop, error) {
	return service.ShopRepository.FindAll()
}

func (service *shopService) Update(shop entities.Shop) {
	service.ShopRepository.Update(shop)
}

func (service *shopService) Delete(shop entities.Shop) {
	service.ShopRepository.Delete(shop)
}

func (service *shopService) GetShopByID(id int) (entities.Shop, error) {
	shop, err := service.ShopRepository.FindByID(id)
	if err != nil {
		return entities.Shop{}, err
	}
	return shop, nil
}
