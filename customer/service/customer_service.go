package service

import (
	"customer/config"
	"customer/dto"
	entity "customer/entity"
	"customer/mapper"
	repository "customer/repository"
	"errors"
)

type CustomerServiceContract interface {
	GetAll() []dto.Customer
	GetByID(uint) dto.Customer

	Create(entity.Customer) (dto.Customer, error)
	Update(entity.Customer) (dto.Customer, error)
	Delete(uint) error
}

type CustomerService struct {
	AddressRepository  repository.AddressRepository
	CustomerRepository repository.CustomerRepository
}

func ProviderCustomerService(
	a repository.AddressRepository,
	c repository.CustomerRepository,
) CustomerService {
	return CustomerService{
		AddressRepository:  a,
		CustomerRepository: c,
	}
}

// Implementation

func (c *CustomerService) GetAll() []dto.Customer {

	// Get All Customer
	customer := c.CustomerRepository.GetAll()

	// Change Customer entity to Customer dto
	return mapper.ToCustomerDtoList(customer)
}

func (c *CustomerService) GetByID(id uint) dto.Customer {

	// Get Customer by Id
	customer := c.CustomerRepository.GetByID(id)

	// Change customer entity to customer dto
	return mapper.ToCustomerDto(customer)
}

func (c *CustomerService) Create(dto dto.Customer) (dto.Customer, error) {

	// Check Phone number first before create data, return error if phone exists
	if !c.CheckPhone(dto) {
		return dto, errors.New(config.PhoneExists)
	}

	// change customer dto to entity to put on database
	customer_entity := mapper.ToCustomerEntity(dto)

	// create customer data
	customer, err := c.CustomerRepository.Create(customer_entity)

	// return error if create customer data error
	if err != nil {
		return dto, err
	}

	// create address data
	err = c.PutAddress(customer.ID, dto.Addresses)

	return mapper.ToCustomerDto(customer), err
}

func (c *CustomerService) Update(dto dto.Customer) (dto.Customer, error) {

	// check customer id first and return error if not exists
	if !c.CheckID(dto.ID) {
		return dto, errors.New(config.CustomerNotFound)
	}

	// map customer dto to entity
	customer_entity := mapper.ToCustomerEntity(dto)

	// update customer data
	customer, err := c.CustomerRepository.Update(customer_entity)

	// map again customer entity to dto
	return mapper.ToCustomerDto(customer), err
}

func (c *CustomerService) Delete(id uint) error {

	// check id first, return errors if not found
	if !c.CheckID(id) {
		return errors.New(config.CustomerNotFound)
	}

	// delete address data first, before deleting customer data
	if err := c.AddressRepository.DeleteByCustomerID(id); err != nil {
		return err
	}

	// delete customer data, and return the result
	return c.CustomerRepository.Delete(id)
}

func (c *CustomerService) CheckPhone(dto dto.Customer) bool {

	phone_number := dto.Phone

	// get customer data by phone number
	customer_data := c.CustomerRepository.GetByPhone(phone_number)

	// return false if data id exists, and if is_active value is true
	if customer_data.ID != 0 && customer_data.IsActive {
		return false
	}

	return true
}

func (c *CustomerService) CheckID(id uint) bool {

	// get customer data by id
	customer_data := c.CustomerRepository.GetByID(id)

	// check if data not found, or check if is_active is false
	if customer_data.ID == 0 || !customer_data.IsActive {
		return false
	}

	return true
}

func (c *CustomerService) PutAddress(id uint, dto []dto.Address) error {

	for _, value := range dto {

		value.CustomerID = id

		// map address dto to entity
		address := mapper.ToAddressEntity(value)

		// call create on address repository to create address value
		if _, err := c.AddressRepository.Create(address); err != nil {
			return err
		}
	}

	return nil
}
