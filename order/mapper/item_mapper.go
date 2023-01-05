package mapper

import (
	"order/dto"
	"order/entity"
)

func ToItemEntity(dto dto.Item) entity.Item {
	return entity.Item{
		ID:          dto.ID,
		OrderID:     dto.OrderID,
		ProductID:   dto.ProductID,
		Quantity:    dto.Quantity,
		Description: dto.Description,
	}
}

func ToItemDto(entity entity.Item, product []entity.Product) dto.Item {

	var dto_item dto.Item

	for _, value := range product {
		if value.ID == entity.ProductID {

			dto_item.ID = entity.ID
			dto_item.OrderID = entity.OrderID
			dto_item.ProductID = entity.ProductID
			dto_item.Name = value.Name
			dto_item.Quantity = entity.Quantity
			dto_item.Description = entity.Description
			dto_item.IsActive = entity.IsActive
		}
	}

	return dto_item
}

func ToItemDtoList(entity []entity.Item, product []entity.Product) []dto.Item {
	item := make([]dto.Item, len(entity))

	for i, value := range entity {
		item[i] = ToItemDto(value, product)
	}

	return item
}
