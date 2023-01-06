package service

import (
	"customer/config"
	"customer/dto"
	entity "customer/entity"
	"customer/mapper"
	repository "customer/repository"
	"errors"
)

type AddressServiceContract interface {
	GetAll() []dto.Address
	GetByID(uint) dto.Address
	GetByCustomerID(uint) dto.Address

	Create(entity.Address) (dto.Address, error)
	Update(entity.Address) (dto.Address, error)
	Delete(uint) error
	DeleteByCustomerID(uint) error
}

type AddressService struct {
	AddressRepository  repository.AddressRepository
	CustomerRepository repository.CustomerRepository
}

func ProviderAddressService(
	a repository.AddressRepository,
	c repository.CustomerRepository,
) AddressService {
	return AddressService{
		AddressRepository:  a,
		CustomerRepository: c,
	}
}

// Implementation

func (a *AddressService) GetAll() []dto.Address {

	// get all address
	address := a.AddressRepository.GetAll()

	// map address entity to address dto
	return mapper.ToAddressDtoList(address)
}

func (a *AddressService) GetByID(id uint) dto.Address {

	// get address by id
	address := a.AddressRepository.GetByID(id)

	// map data from entity to dto
	return mapper.ToAddressDto(address)
}

func (a *AddressService) GetByCustomerID(customer_id uint) []dto.Address {

	// get address by id
	address := a.AddressRepository.GetByCustomerID(customer_id)

	// map address entity to dto
	return mapper.ToAddressDtoList(address)
}

func (a *AddressService) Create(dto dto.Address) (dto.Address, error) {

	// check customer id first before
	if !a.CheckCustomerID(dto.CustomerID) {
		return dto, errors.New(config.CustomerNotFound)
	}

	// Map dto to entity
	address_entity := mapper.ToAddressEntity(dto)

	// Create address data
	address, err := a.AddressRepository.Create(address_entity)

	// Map entity to dto
	return mapper.ToAddressDto(address), err
}

func (a *AddressService) Update(dto dto.Address) (dto.Address, error) {

	// Check ID first, return errors if id not found
	if !a.CheckID(dto.ID) {
		return dto, errors.New(config.AddressNotFound)
	}

	// map dto to entity
	address_entity := mapper.ToAddressEntity(dto)

	// update address
	address, err := a.AddressRepository.Update(address_entity)

	// map entity to dto
	return mapper.ToAddressDto(address), err
}

func (a *AddressService) Delete(id uint) error {

	// check id first, return errors if not found
	if !a.CheckID(id) {
		return errors.New(config.AddressNotFound)
	}

	// delete address
	err := a.AddressRepository.Delete(id)

	// return result
	return err
}

func (a *AddressService) DeleteByCustomerID(customer_id uint) error {

	// check customer id first
	if !a.CheckCustomerID(customer_id) {
		return errors.New(config.CustomerNotFound)
	}

	// delete address by customer id
	err := a.AddressRepository.DeleteByCustomerID(customer_id)

	// return result
	return err
}

func (a *AddressService) CheckID(id uint) bool {

	// get data by id
	address_data := a.AddressRepository.GetByID(id)

	// return false if data not found, or if data is_active is false
	if address_data.ID == 0 || !address_data.IsActive {
		return false
	}

	return true
}

func (a *AddressService) CheckCustomerID(customer_id uint) bool {

	// get customer data by id
	customer_data := a.CustomerRepository.GetByID(customer_id)

	// return false if customer not found, or if customer is not active
	if customer_data.ID == 0 || !customer_data.IsActive {
		return false
	}

	return true
}
