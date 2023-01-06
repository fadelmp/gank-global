package repository

import (
	entity "customer/entity"

	"github.com/jinzhu/gorm"
)

type CustomerRepositoryContract interface {
	GetAll() []entity.Customer
	GetByID(uint) entity.Customer
	GetByPhone(string) entity.Customer

	Create(entity.Customer) (entity.Customer, error)
	Update(entity.Customer) (entity.Customer, error)
	Delete(uint) error
}

type CustomerRepository struct {
	DB *gorm.DB
}

func ProviderCustomerRepository(DB *gorm.DB) CustomerRepository {
	return CustomerRepository{DB: DB}
}

// Implementation

func (c *CustomerRepository) GetAll() []entity.Customer {

	var customers []entity.Customer

	// Get Customer Data with Address Data (has many relationship)
	c.DB.Model(&entity.Customer{}).Preload("Addresses").Find(&customers)

	return customers
}

func (c *CustomerRepository) GetByID(id uint) entity.Customer {

	var customer entity.Customer

	// Get Customer Data by id with Address Data (has many relationship)
	c.DB.Model(&entity.Customer{}).Preload("Addresses").Where("id=?", id).Find(&customer)

	return customer
}

func (c *CustomerRepository) GetByPhone(phone string) entity.Customer {

	var customer entity.Customer

	// Get Customer Data by phone number
	c.DB.Where("phone=?", phone).Find(&customer)

	return customer
}

func (c *CustomerRepository) Create(customer entity.Customer) (entity.Customer, error) {

	// Create Customer Data
	err := c.DB.Create(&customer).Error

	return customer, err
}

func (c *CustomerRepository) Update(customer entity.Customer) (entity.Customer, error) {

	// Update Customer Data by ID
	err := c.DB.Model(&customer).Where("id=?", customer.ID).Update(&customer).Error

	return customer, err
}

func (c *CustomerRepository) Delete(id uint) error {

	var customer entity.Customer

	// Delete Customer Data by change is_active value to false
	err := c.DB.Model(&customer).Where("id=?", id).Updates(map[string]interface{}{
		"is_active": false,
	}).Error

	return err
}
