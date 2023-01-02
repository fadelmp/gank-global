package service

import (
	"errors"
	"product/config"
	"product/dto"
	entity "product/entity"
	"product/mapper"
	repository "product/repository"
)

type CategoryServiceContract interface {
	GetAll() []dto.Category
	GetByID(id uint) dto.Category

	Create(entity.Category) (dto.Category, error)
	Update(entity.Category) (dto.Category, error)
	Delete(id uint) error
}

type CategoryService struct {
	CategoryRepository repository.CategoryRepository
	ProductRepository  repository.ProductRepository
}

func ProviderCategoryService(
	c repository.CategoryRepository,
	p repository.ProductRepository,
) CategoryService {
	return CategoryService{
		CategoryRepository: c,
		ProductRepository:  p,
	}
}

func (c *CategoryService) GetAll() []dto.Category {

	categories := c.CategoryRepository.GetAll()

	return mapper.ToCategoryDtoList(categories)
}

func (c *CategoryService) GetByID(id uint) dto.Category {

	category := c.CategoryRepository.GetByID(id)

	return mapper.ToCategoryDto(category)
}

func (c *CategoryService) Create(dto dto.Category) (dto.Category, error) {

	if !c.CheckName(dto) {
		return dto, errors.New(config.CategoryExists)
	}

	category_entity := mapper.ToCategoryEntity(dto)

	category, err := c.CategoryRepository.Create(category_entity)

	return mapper.ToCategoryDto(category), err
}

func (c *CategoryService) Update(dto dto.Category) (dto.Category, error) {

	if !c.CheckID(dto.ID) {
		return dto, errors.New(config.CategoryNotFound)
	}

	category_entity := mapper.ToCategoryEntity(dto)

	category, err := c.CategoryRepository.Update(category_entity)

	return mapper.ToCategoryDto(category), err
}

func (c *CategoryService) Delete(id uint) error {

	if !c.CheckID(id) {
		return errors.New(config.CategoryNotFound)
	}

	if !c.CheckProduct(id) {
		return errors.New(config.DeleteCategoryFailed)
	}

	return c.CategoryRepository.Delete(id)
}

func (c *CategoryService) CheckName(dto dto.Category) bool {

	category_name := dto.Name

	category_data := c.CategoryRepository.GetByName(category_name)

	if category_data.ID != 0 && category_data.IsActive {
		return false
	}

	return true
}

func (c *CategoryService) CheckID(id uint) bool {

	category_data := c.CategoryRepository.GetByID(id)

	if category_data.ID == 0 || !category_data.IsActive {
		return false
	}

	return true
}

func (c *CategoryService) CheckProduct(id uint) bool {

	product_data := c.ProductRepository.GetByCategoryID(id)

	if len(product_data) > 0 {
		return false
	}

	return true
}
