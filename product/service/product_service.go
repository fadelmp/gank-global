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

func (p *ProductService) GetAll() []dto.Product {

	products := p.ProductRepository.GetAll()

	return mapper.ToProductDtoList(products)
}

func (p *ProductService) GetByID(id uint) dto.Product {

	product := p.ProductRepository.GetByID(id)

	return mapper.ToProductDto(product)
}

func (p *ProductService) GetByCategoryID(category_id uint) []dto.Product {

	products := p.ProductRepository.GetByCategoryID(category_id)

	return mapper.ToProductDtoList(products)
}

func (p *ProductService) Create(dto dto.Product) (dto.Product, error) {

	if !p.CheckName(dto) {
		return dto, errors.New(config.ProductExists)
	}

	if !p.CheckCategory(dto) {
		return dto, errors.New(config.CategoryNotFound)
	}

	product_entity := mapper.ToProductEntity(dto)

	product, err := p.ProductRepository.Create(product_entity)

	return mapper.ToProductDto(product), err
}

func (p *ProductService) Update(dto dto.Product) (dto.Product, error) {

	if !p.CheckID(dto.ID) {
		return dto, errors.New(config.ProductNotFound)
	}

	if !p.CheckCategory(dto) {
		return dto, errors.New(config.CategoryNotFound)
	}

	product_entity := mapper.ToProductEntity(dto)

	product, err := p.ProductRepository.Update(product_entity)

	return mapper.ToProductDto(product), err
}

func (p *ProductService) Delete(id uint) error {

	if !p.CheckID(id) {
		return errors.New(config.ProductNotFound)
	}

	err := p.ProductRepository.Delete(id)

	return err
}

func (p *ProductService) CheckName(dto dto.Product) bool {

	product_name := dto.Name

	product_data := p.ProductRepository.GetByName(product_name)

	if product_data.ID != 0 && product_data.IsActive {
		return false
	}

	return true
}

func (c *ProductService) CheckID(id uint) bool {

	product_data := c.ProductRepository.GetByID(id)

	if product_data.ID == 0 || !product_data.IsActive {
		return false
	}

	return true
}

func (c *ProductService) CheckCategory(dto dto.Product) bool {

	category_id := dto.CategoryID

	category_data := c.CategoryRepository.GetByID(category_id)

	if category_data.ID == 0 || !category_data.IsActive {
		return false
	}

	return true
}
