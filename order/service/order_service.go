package service

import (
	"errors"
	"order/config"
	"order/dto"
	entity "order/entity"
	"order/mapper"
	repository "order/repository"
	"order/request"
	"strconv"
	"time"
)

type OrderServiceContract interface {
	GetAll() ([]dto.Order, error)
	GetByID(id uint) (dto.Order, error)

	Create(entity.Order) error
	Update(entity.Order) error
}

type OrderService struct {
	OrderRepository repository.OrderRepository
	ItemRepository  repository.ItemRepository
	CustomerRequest request.CustomerRequest
	ProductRequest  request.ProductRequest
}

func ProviderOrderService(
	o repository.OrderRepository,
	i repository.ItemRepository,
	c request.CustomerRequest,
	p request.ProductRequest,
) OrderService {
	return OrderService{
		OrderRepository: o,
		ItemRepository:  i,
		CustomerRequest: c,
		ProductRequest:  p,
	}
}

func (o *OrderService) GetAll() ([]dto.Order, error) {

	orders := o.OrderRepository.GetAll()

	customers, err := o.CustomerRequest.GetAllCustomer()

	products, err := o.ProductRequest.GetAllProduct()

	return mapper.ToOrderDtoList(orders, customers, products), err
}

func (o *OrderService) GetByID(id uint) (dto.Order, error) {

	order := o.OrderRepository.GetByID(id)

	customer, err := o.CustomerRequest.GetCustomerById(order.CustomerID)

	products, err := o.ProductRequest.GetAllProduct()

	return mapper.ToOrderDto(order, customer, products), err
}

func (o *OrderService) Create(dto dto.Order) error {

	dto.Number = o.GenerateOrderNumber()

	order_entity := mapper.ToOrderEntity(dto)

	order, err := o.OrderRepository.Create(order_entity)

	if err != nil {
		return err
	}

	err = o.PutItem(order.ID, dto.Item)

	return err
}

func (o *OrderService) Update(dto dto.Order) error {

	if !o.CheckID(dto.ID) {
		return errors.New(config.OrderNotFound)
	}

	Order_entity := mapper.ToOrderEntity(dto)

	_, err := o.OrderRepository.Update(Order_entity)

	return err
}

func (o *OrderService) CheckID(id uint) bool {

	order_data := o.OrderRepository.GetByID(id)

	if order_data.ID == 0 {
		return false
	}

	return true
}

func (o *OrderService) PutItem(id uint, dto []dto.Item) error {

	for _, value := range dto {

		value.OrderID = id
		address := mapper.ToItemEntity(value)

		if _, err := o.ItemRepository.Create(address); err != nil {
			return err
		}
	}

	return nil
}

func (o *OrderService) GenerateOrderNumber() string {

	t := time.Now()
	return strconv.FormatInt(t.UnixNano(), 10)

}
