package routes

import (
	"product/controller"

	"github.com/labstack/echo"
)

func ProductRoutes(routes *echo.Echo, api controller.ProductController) {

	product := routes.Group("/product")
	{
		product.GET("", api.GetAll)
		product.GET("/:id", api.GetByID)
		product.GET("/category/:category_id", api.GetByCategoryID)

		product.POST("", api.Create)
		product.PUT("", api.Update)
		product.DELETE("/:id", api.Delete)
	}
}
