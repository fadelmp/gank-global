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
	GetByID(id uint) dto.Address
	GetByCustomerID(customer_id uint) dto.Address

	Create(entity.Address) (dto.Address, error)
	Update(entity.Address) (dto.Address, error)
	Delete(id uint) error
	DeleteByCustomerID(customer_id uint) error
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

func (a *AddressService) GetAll() []dto.Address {

	address := a.AddressRepository.GetAll()

	return mapper.ToAddressDtoList(address)
}

func (a *AddressService) GetByID(id uint) dto.Address {

	address := a.AddressRepository.GetByID(id)

	return mapper.ToAddressDto(address)
}

func (a *AddressService) GetByCustomerID(customer_id uint) []dto.Address {

	address := a.AddressRepository.GetByCustomerID(customer_id)

	return mapper.ToAddressDtoList(address)
}

func (a *AddressService) Create(dto dto.Address) (dto.Address, error) {

	if !a.CheckCustomerID(dto.CustomerID) {
		return dto, errors.New(config.CustomerNotFound)
	}

	address_entity := mapper.ToAddressEntity(dto)

	address, err := a.AddressRepository.Create(address_entity)

	return mapper.ToAddressDto(address), err
}

func (a *AddressService) Update(dto dto.Address) (dto.Address, error) {

	if !a.CheckID(dto.ID) {
		return dto, errors.New(config.AddressNotFound)
	}

	address_entity := mapper.ToAddressEntity(dto)

	address, err := a.AddressRepository.Update(address_entity)

	return mapper.ToAddressDto(address), err
}

func (a *AddressService) Delete(id uint) error {

	if !a.CheckID(id) {
		return errors.New(config.AddressNotFound)
	}

	err := a.AddressRepository.Delete(id)

	return err
}

func (a *AddressService) DeleteByCustomerID(customer_id uint) error {

	if !a.CheckCustomerID(customer_id) {
		return errors.New(config.CustomerNotFound)
	}

	err := a.AddressRepository.DeleteByCustomerID(customer_id)

	return err
}

func (a *AddressService) CheckID(id uint) bool {

	address_data := a.AddressRepository.GetByID(id)

	if address_data.ID == 0 {
		return false
	}

	return true
}

func (a *AddressService) CheckCustomerID(customer_id uint) bool {

	customer_data := a.CustomerRepository.GetByID(customer_id)

	if customer_data.ID == 0 {
		return false
	}

	return true
}
