package mapper

import (
	"order/dto"
	"order/entity"
)

func ToOrderEntity(dto dto.Order) entity.Order {
	return entity.Order{
		ID:          dto.ID,
		Number:      dto.Number,
		StatusID:    dto.StatusID,
		Total:       dto.Total,
		Date:        dto.Date,
		Description: dto.Description,
	}
}

func ToOrderDto(
	order entity.Order,
	customer entity.Customer,
	product []entity.Product,
) dto.Order {

	var order_dto dto.Order
	for _, value := range customer.Address {

		if value.ID == order.AddressID {

			order_dto.ID = order.ID
			order_dto.Number = order.Number
			order_dto.StatusID = order.StatusID
			order_dto.CustomerName = customer.Name
			order_dto.CustomerPhone = customer.Phone
			order_dto.CustomerEmail = customer.Email
			order_dto.AddressLine = value.AddressLine
			order_dto.City = value.City
			order_dto.Province = value.Province
			order_dto.Country = value.Country
			order_dto.PostalCode = value.PostalCode
			order_dto.Total = order.Total
			order_dto.Date = order.Date
			order_dto.Description = order.Description
			order_dto.Item = ToItemDtoList(order.Items, product)
		}
	}

	return order_dto
}

func ToOrderDtoList(
	entity []entity.Order,
	customer []entity.Customer,
	product []entity.Product,
) []dto.Order {

	order := make([]dto.Order, len(entity))

	for i, order_value := range entity {

		for _, customer_value := range customer {

			if customer_value.ID == order_value.CustomerID {
				order[i] = ToOrderDto(order_value, customer_value, product)
			}
		}
	}

	return order
}
