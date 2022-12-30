package routes

import (
	"customer/controller"

	"github.com/labstack/echo"
)

func AddressRoutes(routes *echo.Echo, api controller.AddressController) {

	address := routes.Group("/address")
	{
		address.GET("", api.GetAll)
		address.GET("/:id", api.GetByID)
		address.GET("/customer/:customer_id", api.GetByCustomerID)

		address.POST("", api.Create)
		address.PUT("", api.Update)
		address.DELETE("/:id", api.Delete)
	}
}
