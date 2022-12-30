package controller

import (
	"customer/config"
	"customer/dto"
	"customer/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type CustomerController struct {
	CustomerService service.CustomerService
}

func ProviderCustomerController(c service.CustomerService) CustomerController {
	return CustomerController{
		CustomerService: c,
	}
}

func (c *CustomerController) GetAll(e echo.Context) error {

	customers := c.CustomerService.GetAll()

	if len(customers) == 0 {
		return config.SuccessResponse(e, http.StatusNoContent, config.CustomerNotFound)
	}

	return config.SuccessResponse(e, http.StatusOK, customers)
}

func (c *CustomerController) GetByID(e echo.Context) error {

	id, _ := strconv.ParseUint(e.Param("id"), 10, 32)

	customer := c.CustomerService.GetByID(uint(id))

	if customer.ID == 0 {
		return config.SuccessResponse(e, http.StatusNoContent, config.CustomerNotFound)
	}

	return config.SuccessResponse(e, http.StatusOK, customer)
}

func (c *CustomerController) Create(e echo.Context) error {

	var customer dto.Customer

	if e.Bind(&customer) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	res, err := c.CustomerService.Create(customer)

	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, http.StatusOK, res)
}

func (c *CustomerController) Update(e echo.Context) error {

	var customer dto.Customer

	if e.Bind(&customer) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	res, err := c.CustomerService.Update(customer)

	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, http.StatusOK, res)
}

func (a *CustomerController) Delete(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	err = a.CustomerService.Delete(uint(id))
	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, http.StatusOK, "Delete Success")
}
