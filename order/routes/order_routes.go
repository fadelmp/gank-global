package routes

import (
	"order/controller"

	"github.com/labstack/echo"
)

func OrderRoutes(routes *echo.Echo, api controller.OrderController) {

	order := routes.Group("/order")
	{
		order.GET("", api.GetAll)
		order.GET("/:id", api.GetByID)

		order.POST("", api.Create)
	}
}
