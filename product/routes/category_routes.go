package routes

import (
	"product/controller"

	"github.com/labstack/echo"
)

func CategoryRoutes(routes *echo.Echo, api controller.CategoryController) {

	category := routes.Group("/category")
	{
		category.GET("", api.GetAll)
		category.GET("/:id", api.GetByID)

		category.POST("", api.Create)
		category.PUT("", api.Update)
		category.DELETE("/:id", api.Delete)
	}
}
