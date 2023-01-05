package repository

import (
	entity "order/entity"
	"time"

	"github.com/jinzhu/gorm"
)

type OrderRepositoryContract interface {
	GetAll() []entity.Order
	GetByID(uint) entity.Order
	GetByCustomerID(uint) []entity.Order
	GetByStatusID(uint) []entity.Order
	GetByDate(time.Time, time.Time) []entity.Order

	Create(entity.Order) (entity.Order, error)
	Update(entity.Order) (entity.Order, error)
}

type OrderRepository struct {
	DB *gorm.DB
}

func ProviderOrderRepository(DB *gorm.DB) OrderRepository {
	return OrderRepository{DB: DB}
}

func (o *OrderRepository) GetAll() []entity.Order {

	var orders []entity.Order

	o.DB.Model(&entity.Order{}).Preload("Items").Find(&orders)

	return orders
}

func (o *OrderRepository) GetByID(id uint) entity.Order {

	var order entity.Order

	o.DB.Model(&entity.Order{}).Preload("Items").Where("id=?", id).Find(&order)

	return order
}

func (o *OrderRepository) GetByCustomerID(customer_id uint) []entity.Order {

	var orders []entity.Order

	o.DB.Model(&entity.Order{}).Preload("Items").Where("customer_id=?", customer_id).Find(&orders)

	return orders
}

func (o *OrderRepository) GetByStatusID(status_id uint) []entity.Order {

	var orders []entity.Order

	o.DB.Where("id=?", status_id).Find(&orders)

	return orders
}

func (o *OrderRepository) Create(order entity.Order) (entity.Order, error) {

	err := o.DB.Create(&order).Error

	return order, err
}

func (o *OrderRepository) Update(order entity.Order) (entity.Order, error) {

	err := o.DB.Model(&order).Where("id=?", order.ID).Update(&order).Error

	return order, err
}
