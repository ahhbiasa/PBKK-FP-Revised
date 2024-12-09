package service

import (
	entities "PBKK-FP-Revised/entities"
	"PBKK-FP-Revised/repository"
)

type CategoryService interface {
	Save(entities.Category) entities.Category
	Update(category entities.Category)
	Delete(category entities.Category)
	FindAll() []entities.Category
	GetCategoryByID(id int) (entities.Category, error)
}

type categoryService struct {
	CategoryryRepository repository.CategoryRepository
}

func New(repo repository.CategoryRepository) CategoryService {
	return &categoryService{
		CategoryryRepository: repo,
	}
}

func (service *categoryService) Save(category entities.Category) entities.Category {
	service.CategoryryRepository.Save(category)
	return category
}

func (service *categoryService) FindAll() []entities.Category {
	return service.CategoryryRepository.FindAll()
}

func (service *categoryService) Update(category entities.Category) {
	service.CategoryryRepository.Update(category)
}

func (service *categoryService) Delete(category entities.Category) {
	service.CategoryryRepository.Delete(category)
}

func (service *categoryService) GetCategoryByID(id int) (entities.Category, error) {
	return service.CategoryryRepository.FindByID(id)
}
