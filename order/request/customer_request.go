package request

import (
	"encoding/json"
	entity "order/entity"
	"os"
	"strconv"
)

type CustomerRequestContract interface {
	GetAllCustomer() ([]entity.Customer, error)
	GetCustomerById(uint) (entity.Customer, error)
}

type CustomerRequest struct{}

func ProviderCustomerRequest() CustomerRequest {
	return CustomerRequest{}
}

func (c *CustomerRequest) GetAllCustomer() ([]entity.Customer, error) {

	var customer []entity.Customer
	var customer_response entity.CustomerAllResponse

	var uri = os.Getenv("CUSTOMER_SERVICE_URI")
	uri += "customer"

	result, err := GetRequest(uri)
	if err == nil {
		json.Unmarshal([]byte(result), &customer_response)
		customer = customer_response.Data
	}

	return customer, err

}

func (c *CustomerRequest) GetCustomerById(customer_id uint) (entity.Customer, error) {

	var customer entity.Customer
	var customer_response entity.CustomerResponse

	var uri = os.Getenv("CUSTOMER_SERVICE_URI")
	uri += "customer/"
	uri += strconv.FormatUint(uint64(customer_id), 10)

	result, err := GetRequest(uri)
	if err == nil {
		json.Unmarshal([]byte(result), &customer_response)
		customer = customer_response.Data
	}

	return customer, err

}
