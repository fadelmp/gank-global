package injection

import (
	"customer/controller"
	repository "customer/repository"
	"customer/service"

	"github.com/jinzhu/gorm"
)

func AddressInjection(db *gorm.DB) controller.AddressController {

	AddressRepository := repository.ProviderAddressRepository(db)
	CustomerRepository := repository.ProviderCustomerRepository(db)

	AddressService := service.ProviderAddressService(AddressRepository, CustomerRepository)
	AddressController := controller.ProviderAddressController(AddressService)

	return AddressController
}
