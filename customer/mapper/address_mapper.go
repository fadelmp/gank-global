package mapper

import (
	"customer/dto"
	"customer/entity"
)

func ToAddressEntity(dto dto.Address) entity.Address {
	return entity.Address{
		ID:          dto.ID,
		CustomerID:  dto.CustomerID,
		AddressLine: dto.AddressLine,
		City:        dto.City,
		Province:    dto.Province,
		Country:     dto.Country,
		PostalCode:  dto.PostalCode,
		IsActive:    dto.IsActive,
	}
}

func ToAddressDto(entity entity.Address) dto.Address {
	return dto.Address{
		ID:          entity.ID,
		CustomerID:  entity.CustomerID,
		AddressLine: entity.AddressLine,
		City:        entity.City,
		Province:    entity.Province,
		Country:     entity.Country,
		PostalCode:  entity.PostalCode,
		IsActive:    entity.IsActive,
	}
}

func ToAddressDtoList(entity []entity.Address) []dto.Address {
	address := make([]dto.Address, len(entity))

	for i, value := range entity {
		address[i] = ToAddressDto(value)
	}

	return address
}
