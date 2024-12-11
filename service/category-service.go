package service

import (
	entities "PBKK-FP-Revised/entities"
	"PBKK-FP-Revised/repository"
)

type CategoryService interface {
	Save(category entities.Category) error
	Update(category entities.Category)
	Delete(category entities.Category)
	FindAll() ([]entities.Category, error) // Include error in the return value
	GetCategoryByID(id int) (entities.Category, error)
}

type categoryService struct {
	CategoryryRepository repository.CategoryRepository
}

func NewCategoryController(repo repository.CategoryRepository) CategoryService {
	return &categoryService{
		CategoryryRepository: repo,
	}
}

func (service *categoryService) Save(category entities.Category) error {
	err := service.CategoryryRepository.Save(category)
	if err != nil {
		return err // Return an empty category and the error
	}
	return nil // Return the saved category and nil error
}

func (service *categoryService) FindAll() ([]entities.Category, error) {
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
