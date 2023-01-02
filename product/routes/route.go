package routes

import (
	"product/injection"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func Init(routes *echo.Echo, db *gorm.DB) *echo.Echo {

	// Category Route & Injection
	category := injection.CategoryInjection(db)
	CategoryRoutes(routes, category)

	// Product Route & Injection
	product := injection.ProductInjection(db)
	ProductRoutes(routes, product)

	return routes
}
