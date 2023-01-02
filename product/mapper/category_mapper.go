package mapper

import (
	"product/dto"
	"product/entity"
)

func ToCategoryEntity(dto dto.Category) entity.Category {
	return entity.Category{
		ID:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
	}
}

func ToCategoryDto(entity entity.Category) dto.Category {
	return dto.Category{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		IsActive:    entity.IsActive,
	}
}

func ToCategoryDtoList(entity []entity.Category) []dto.Category {
	category := make([]dto.Category, len(entity))

	for i, value := range entity {
		category[i] = ToCategoryDto(value)
	}

	return category
}
