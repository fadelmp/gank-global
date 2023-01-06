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

// Implementation

func (c *CategoryService) GetAll() []dto.Category {

	// get all category
	categories := c.CategoryRepository.GetAll()

	// map category entity to dto
	return mapper.ToCategoryDtoList(categories)
}

func (c *CategoryService) GetByID(id uint) dto.Category {

	// get all category
	category := c.CategoryRepository.GetByID(id)

	// map category entity to dto
	return mapper.ToCategoryDto(category)
}

func (c *CategoryService) Create(dto dto.Category) (dto.Category, error) {

	// check category name first, if name exists then return errors
	if !c.CheckName(dto) {
		return dto, errors.New(config.CategoryExists)
	}

	// map dto to entity
	category_entity := mapper.ToCategoryEntity(dto)

	// create category
	category, err := c.CategoryRepository.Create(category_entity)

	// map category entity to dto and return
	return mapper.ToCategoryDto(category), err
}

func (c *CategoryService) Update(dto dto.Category) (dto.Category, error) {

	// check id first, return errors if not found
	if !c.CheckID(dto.ID) {
		return dto, errors.New(config.CategoryNotFound)
	}

	// map dto to entity
	category_entity := mapper.ToCategoryEntity(dto)

	// update category
	category, err := c.CategoryRepository.Update(category_entity)

	// return map entity to dto category
	return mapper.ToCategoryDto(category), err
}

func (c *CategoryService) Delete(id uint) error {

	// check id first, return errors if id not found
	if !c.CheckID(id) {
		return errors.New(config.CategoryNotFound)
	}

	// check product first, return errors if there are product that use this category
	if !c.CheckProduct(id) {
		return errors.New(config.DeleteCategoryFailed)
	}

	// delete category
	return c.CategoryRepository.Delete(id)
}

func (c *CategoryService) CheckName(dto dto.Category) bool {

	category_name := dto.Name

	// get category by name
	category_data := c.CategoryRepository.GetByName(category_name)

	// return error if category exists
	if category_data.ID != 0 && category_data.IsActive {
		return false
	}

	return true
}

func (c *CategoryService) CheckID(id uint) bool {

	// get category by id
	category_data := c.CategoryRepository.GetByID(id)

	// return false if category is not exists
	if category_data.ID == 0 || !category_data.IsActive {
		return false
	}

	return true
}

func (c *CategoryService) CheckProduct(id uint) bool {

	// get product by category
	product_data := c.ProductRepository.GetByCategoryID(id)

	// return false if product exists
	if len(product_data) > 0 {
		return false
	}

	return true
}
