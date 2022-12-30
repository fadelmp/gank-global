package controller

import (
	"customer/config"
	"customer/dto"
	"customer/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
)

type AddressController struct {
	AddressService service.AddressService
}

func ProviderAddressController(a service.AddressService) AddressController {
	return AddressController{AddressService: a}
}

func (a *AddressController) GetAll(e echo.Context) error {

	addresses := a.AddressService.GetAll()

	if len(addresses) == 0 {
		return config.SuccessResponse(e, http.StatusNoContent, config.CustomerNotFound)
	}

	return config.SuccessResponse(e, http.StatusOK, addresses)
}

func (a *AddressController) GetByID(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)
	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	address := a.AddressService.GetByID(uint(id))
	if address.ID == 0 {
		return config.SuccessResponse(e, http.StatusNoContent, config.CustomerNotFound)
	}

	return config.SuccessResponse(e, http.StatusOK, address)
}

func (a *AddressController) GetByCustomerID(e echo.Context) error {

	customer_id, err := strconv.ParseUint(e.Param("customer_id"), 10, 64)
	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	addresses := a.AddressService.GetByCustomerID(uint(customer_id))
	if len(addresses) == 0 {
		return config.SuccessResponse(e, http.StatusNoContent, addresses)
	}

	return config.SuccessResponse(e, http.StatusOK, addresses)
}

func (a *AddressController) Create(e echo.Context) error {

	var address dto.Address
	if e.Bind(&address) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	res, err := a.AddressService.Create(address)
	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, http.StatusOK, res)
}

func (a *AddressController) Update(e echo.Context) error {

	var address dto.Address
	if e.Bind(&address) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	res, err := a.AddressService.Update(address)
	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, http.StatusOK, res)
}

func (a *AddressController) Delete(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 32)
	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	err = a.AddressService.Delete(uint(id))
	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, http.StatusOK, "Delete Success")
}
