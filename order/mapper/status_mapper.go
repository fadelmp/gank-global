package mapper

import (
	"order/dto"
	"order/entity"
)

func ToStatusEntity(dto dto.Status) entity.Status {
	return entity.Status{
		ID:               dto.ID,
		Name:             dto.Name,
		FrontDescription: dto.FrontDescription,
		BackDescription:  dto.BackDescription,
	}
}

func ToStatusDto(entity entity.Status) dto.Status {
	return dto.Status{
		ID:               entity.ID,
		Name:             entity.Name,
		FrontDescription: entity.FrontDescription,
		BackDescription:  entity.BackDescription,
		IsActive:         entity.IsActive,
	}
}

func ToStatusDtoList(entity []entity.Status) []dto.Status {
	statuses := make([]dto.Status, len(entity))

	for i, value := range entity {
		statuses[i] = ToStatusDto(value)
	}

	return statuses
}
