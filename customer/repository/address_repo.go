package repository

import (
	entity "customer/entity"

	"github.com/jinzhu/gorm"
)

type AddressRepositoryContract interface {
	GetAll() []entity.Address
	GetByID(uint) entity.Address
	GetByCustomerID(uint) []entity.Address

	Create(entity.Address) (entity.Address, error)
	Update(entity.Address) (entity.Address, error)
	Delete(uint) error
	DeleteByCustomerID(uint) error
}

type AddressRepository struct {
	DB *gorm.DB
}

func ProviderAddressRepository(DB *gorm.DB) AddressRepository {
	return AddressRepository{DB: DB}
}

func (a *AddressRepository) GetAll() []entity.Address {

	var addresses []entity.Address

	a.DB.Find(&addresses)

	return addresses
}

func (a *AddressRepository) GetByID(id uint) entity.Address {

	var address entity.Address

	a.DB.Where("id=?", id).Find(&address)

	return address
}

func (a *AddressRepository) GetByCustomerID(customer_id uint) []entity.Address {

	var addresses []entity.Address

	a.DB.Where("customer_id=?", customer_id).Find(&addresses)

	return addresses

}

func (a *AddressRepository) Create(address entity.Address) (entity.Address, error) {

	err := a.DB.Create(&address).Error

	return address, err
}

func (a *AddressRepository) Update(address entity.Address) (entity.Address, error) {

	err := a.DB.Model(&address).Where("id=?", address.ID).Update(&address).Error

	return address, err
}

func (a *AddressRepository) Delete(id uint) error {

	var address entity.Address

	err := a.DB.Model(&address).Where("id=?", id).Updates(map[string]interface{}{
		"is_active": false,
	}).Error

	return err
}

func (a *AddressRepository) DeleteByCustomerID(customer_id uint) error {

	var address entity.Address

	err := a.DB.Model(&address).Where("customer_id=?", customer_id).Updates(map[string]interface{}{
		"is_active": false,
	}).Error

	return err
}
