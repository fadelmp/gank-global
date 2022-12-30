package routes

import (
	"customer/controller"

	"github.com/labstack/echo"
)

func CustomerRoutes(routes *echo.Echo, api controller.CustomerController) {

	customer := routes.Group("/customer")
	{
		customer.GET("", api.GetAll)
		customer.GET("/:id", api.GetByID)

		customer.POST("", api.Create)
		customer.PUT("", api.Update)
		customer.DELETE("/:id", api.Delete)
	}
}
