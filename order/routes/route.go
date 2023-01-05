package routes

import (
	"order/injection"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func Init(routes *echo.Echo, db *gorm.DB) *echo.Echo {

	// Status Route & Injection
	status := injection.StatusInjection(db)
	StatusRoutes(routes, status)

	// Order Route & Injection
	order := injection.OrderInjection(db)
	OrderRoutes(routes, order)
	return routes
}
