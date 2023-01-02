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

func (c *CustomerService) GetAll() []dto.Customer {

	customer := c.CustomerRepository.GetAll()

	return mapper.ToCustomerDtoList(customer)
}

func (c *CustomerService) GetByID(id uint) dto.Customer {

	customer := c.CustomerRepository.GetByID(id)

	return mapper.ToCustomerDto(customer)
}

func (c *CustomerService) Create(dto dto.Customer) (dto.Customer, error) {

	if !c.CheckPhone(dto) {
		return dto, errors.New(config.PhoneExists)
	}

	customer_entity := mapper.ToCustomerEntity(dto)

	customer, err := c.CustomerRepository.Create(customer_entity)

	if err != nil {
		return dto, err
	}

	err = c.PutAddress(customer.ID, dto.Addresses)

	return mapper.ToCustomerDto(customer), err
}

func (c *CustomerService) Update(dto dto.Customer) (dto.Customer, error) {

	if !c.CheckID(dto.ID) {
		return dto, errors.New(config.CustomerNotFound)
	}

	customer_entity := mapper.ToCustomerEntity(dto)

	customer, err := c.CustomerRepository.Update(customer_entity)

	return mapper.ToCustomerDto(customer), err
}

func (c *CustomerService) Delete(id uint) error {

	if !c.CheckID(id) {
		return errors.New(config.CustomerNotFound)
	}

	if err := c.AddressRepository.DeleteByCustomerID(id); err != nil {
		return err
	}

	return c.CustomerRepository.Delete(id)
}

func (c *CustomerService) CheckPhone(dto dto.Customer) bool {

	phone_number := dto.Phone

	customer_data := c.CustomerRepository.GetByPhone(phone_number)

	if customer_data.ID != 0 && customer_data.IsActive {
		return false
	}

	return true
}

func (c *CustomerService) CheckID(id uint) bool {

	customer_data := c.CustomerRepository.GetByID(id)

	if customer_data.ID == 0 || !customer_data.IsActive {
		return false
	}

	return true
}

func (c *CustomerService) PutAddress(id uint, dto []dto.Address) error {

	for _, value := range dto {

		value.CustomerID = id
		address := mapper.ToAddressEntity(value)

		if _, err := c.AddressRepository.Create(address); err != nil {
			return err
		}
	}

	return nil
}
