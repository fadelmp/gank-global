package mapper

import (
	"product/dto"
	"product/entity"
)

func ToProductEntity(dto dto.Product) entity.Product {
	return entity.Product{
		ID:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
		CategoryID:  dto.CategoryID,
		Price:       dto.Price,
		Stock:       dto.Stock,
	}
}

func ToProductDto(entity entity.Product) dto.Product {
	return dto.Product{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		CategoryID:  entity.CategoryID,
		Price:       entity.Price,
		Stock:       entity.Stock,
		IsActive:    entity.IsActive,
	}
}

func ToProductDtoList(entity []entity.Product) []dto.Product {
	product := make([]dto.Product, len(entity))

	for i, value := range entity {
		product[i] = ToProductDto(value)
	}

	return product
}
