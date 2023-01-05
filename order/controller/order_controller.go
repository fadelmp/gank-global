package controller

import (
	"net/http"
	"order/config"
	"order/dto"
	"order/service"
	"strconv"

	"github.com/labstack/echo"
)

type OrderController struct {
	OrderService service.OrderService
}

func ProviderOrderController(o service.OrderService) OrderController {
	return OrderController{
		OrderService: o,
	}
}

func (o *OrderController) GetAll(e echo.Context) error {

	orders, err := o.OrderService.GetAll()

	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.GetOrderFailed)
	}

	if len(orders) == 0 {
		return config.SuccessResponse(e, http.StatusNoContent, config.StatusNotFound)
	}

	return config.SuccessResponse(e, http.StatusOK, orders)
}

func (o *OrderController) GetByID(e echo.Context) error {

	id, _ := strconv.ParseUint(e.Param("id"), 10, 32)

	order, err := o.OrderService.GetByID(uint(id))

	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.GetOrderFailed)
	}

	if order.ID == 0 {
		return config.SuccessResponse(e, http.StatusNoContent, config.StatusNotFound)
	}

	return config.SuccessResponse(e, http.StatusOK, order)
}

func (o *OrderController) Create(e echo.Context) error {

	var order dto.Order

	if e.Bind(&order) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	res, err := o.OrderService.Create(order)

	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, http.StatusOK, res)
}
