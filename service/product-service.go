package service

import (
	entities "PBKK-FP-Revised/entities"
	"PBKK-FP-Revised/repository"
)

type ProductService interface {
	Save(product entities.Product) error
	Update(product entities.Product)
	Delete(product entities.Product)
	FindAll() ([]entities.Product, error) // Include error in the return value
	GetProductByID(id int) (entities.Product, error)
	GetCategoryByID(id int) (entities.Category, error) // Add this method
	GetShopByID(id int) (entities.Shop, error)
}

type productService struct {
	ProductRepository  repository.ProductRepository
	CategoryRepository repository.CategoryRepository
	ShopRepository     repository.ShopRepository
}

func NewProductService(productrepo repository.ProductRepository, categoryrepo repository.CategoryRepository, shoprepo repository.ShopRepository) ProductService {
	return &productService{
		ProductRepository:  productrepo,
		CategoryRepository: categoryrepo,
		ShopRepository:     shoprepo,
	}
}

func (service *productService) Save(product entities.Product) error {
	err := service.ProductRepository.Save(product)
	if err != nil {
		return err
	}
	return nil
}

func (service *productService) FindAll() ([]entities.Product, error) {
	return service.ProductRepository.FindAll()
}

func (service *productService) Update(product entities.Product) {
	service.ProductRepository.Update(product)
}

func (service *productService) Delete(product entities.Product) {
	service.ProductRepository.Delete(product)
}

func (service *productService) GetProductByID(id int) (entities.Product, error) {
	product, err := service.ProductRepository.FindByID(id)
	if err != nil {
		return entities.Product{}, err
	}
	return product, nil
}

func (service *productService) GetCategoryByID(id int) (entities.Category, error) {
	category, err := service.CategoryRepository.FindByID(id)
	if err != nil {
		return entities.Category{}, err
	}
	return category, nil
}

func (service *productService) GetShopByID(id int) (entities.Shop, error) {
	shop, err := service.ShopRepository.FindByID(id)
	if err != nil {
		return entities.Shop{}, err
	}
	return shop, nil
}
