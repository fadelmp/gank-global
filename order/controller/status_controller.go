package controller

import (
	"net/http"
	"order/config"
	"order/dto"
	"order/service"
	"strconv"

	"github.com/labstack/echo"
)

type StatusController struct {
	StatusService service.StatusService
}

func ProviderStatusController(c service.StatusService) StatusController {
	return StatusController{
		StatusService: c,
	}
}

func (c *StatusController) GetAll(e echo.Context) error {

	statuses := c.StatusService.GetAll()

	if len(statuses) == 0 {
		return config.SuccessResponse(e, http.StatusNoContent, config.StatusNotFound)
	}

	return config.SuccessResponse(e, http.StatusOK, statuses)
}

func (c *StatusController) GetByID(e echo.Context) error {

	id, _ := strconv.ParseUint(e.Param("id"), 10, 32)

	status := c.StatusService.GetByID(uint(id))

	if status.ID == 0 {
		return config.SuccessResponse(e, http.StatusNoContent, config.StatusNotFound)
	}

	return config.SuccessResponse(e, http.StatusOK, status)
}

func (c *StatusController) Create(e echo.Context) error {

	var status dto.Status

	if e.Bind(&status) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	res, err := c.StatusService.Create(status)

	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, http.StatusOK, res)
}

func (c *StatusController) Update(e echo.Context) error {

	var status dto.Status

	if e.Bind(&status) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	res, err := c.StatusService.Update(status)

	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, http.StatusOK, res)
}

func (a *StatusController) Delete(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	err = a.StatusService.Delete(uint(id))
	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, http.StatusOK, "Delete Success")
}
