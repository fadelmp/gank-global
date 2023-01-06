package service

import (
	"errors"
	"product/config"
	"product/dto"
	entity "product/entity"
	"product/mapper"
	repository "product/repository"
)

type ProductServiceContract interface {
	GetAll() []dto.Product
	GetByID(id uint) dto.Product
	GetByCategoryID(category_id uint) []dto.Product

	Create(entity.Product) (dto.Product, error)
	Update(entity.Product) (dto.Product, error)
	Delete(id uint) error
}

type ProductService struct {
	CategoryRepository repository.CategoryRepository
	ProductRepository  repository.ProductRepository
}

func ProviderProductService(
	c repository.CategoryRepository,
	p repository.ProductRepository,
) ProductService {
	return ProductService{
		CategoryRepository: c,
		ProductRepository:  p,
	}
}

// Implementation

func (p *ProductService) GetAll() []dto.Product {

	// get all product
	products := p.ProductRepository.GetAll()

	// map product entity to dto
	return mapper.ToProductDtoList(products)
}

func (p *ProductService) GetByID(id uint) dto.Product {

	// get product by id
	product := p.ProductRepository.GetByID(id)

	// map product entity to dto
	return mapper.ToProductDto(product)
}

func (p *ProductService) GetByCategoryID(category_id uint) []dto.Product {

	// get product by category id
	products := p.ProductRepository.GetByCategoryID(category_id)

	// map product dto to entity
	return mapper.ToProductDtoList(products)
}

func (p *ProductService) Create(dto dto.Product) (dto.Product, error) {

	// check product name, return errors if exists
	if !p.CheckName(dto) {
		return dto, errors.New(config.ProductExists)
	}

	// check category first, return errors if not found
	if !p.CheckCategory(dto) {
		return dto, errors.New(config.CategoryNotFound)
	}

	// map product dto to entity
	product_entity := mapper.ToProductEntity(dto)

	// create product
	product, err := p.ProductRepository.Create(product_entity)

	// return map entity to dto
	return mapper.ToProductDto(product), err
}

func (p *ProductService) Update(dto dto.Product) (dto.Product, error) {

	// check product id, return errors if not found
	if !p.CheckID(dto.ID) {
		return dto, errors.New(config.ProductNotFound)
	}

	// check product category, return errors if not found
	if !p.CheckCategory(dto) {
		return dto, errors.New(config.CategoryNotFound)
	}

	// map product dto to entity
	product_entity := mapper.ToProductEntity(dto)

	// update product
	product, err := p.ProductRepository.Update(product_entity)

	// return map entity to dto
	return mapper.ToProductDto(product), err
}

func (p *ProductService) Delete(id uint) error {

	// check id first, return errors if product not found
	if !p.CheckID(id) {
		return errors.New(config.ProductNotFound)
	}

	// delete product
	err := p.ProductRepository.Delete(id)

	return err
}

func (p *ProductService) CheckName(dto dto.Product) bool {

	product_name := dto.Name

	// get product by name
	product_data := p.ProductRepository.GetByName(product_name)

	// return false if product exists
	if product_data.ID != 0 && product_data.IsActive {
		return false
	}

	return true
}

func (c *ProductService) CheckID(id uint) bool {

	// get product by id
	product_data := c.ProductRepository.GetByID(id)

	// return false if product not found
	if product_data.ID == 0 || !product_data.IsActive {
		return false
	}

	return true
}

func (c *ProductService) CheckCategory(dto dto.Product) bool {

	category_id := dto.CategoryID

	// get category by category id
	category_data := c.CategoryRepository.GetByID(category_id)

	// return false if category not found
	if category_data.ID == 0 || !category_data.IsActive {
		return false
	}

	return true
}
