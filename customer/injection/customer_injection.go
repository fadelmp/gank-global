package injection

import (
	"customer/controller"
	repository "customer/repository"
	"customer/service"

	"github.com/jinzhu/gorm"
)

func CustomerInjection(db *gorm.DB) controller.CustomerController {

	AddressRepository := repository.ProviderAddressRepository(db)
	CustomerRepository := repository.ProviderCustomerRepository(db)

	CustomerService := service.ProviderCustomerService(AddressRepository, CustomerRepository)
	CustomerController := controller.ProviderCustomerController(CustomerService)

	return CustomerController
}
