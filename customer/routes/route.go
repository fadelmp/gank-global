package routes

import (
	"customer/injection"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func Init(routes *echo.Echo, db *gorm.DB) *echo.Echo {

	// Customer Route & Injection
	customer := injection.CustomerInjection(db)
	CustomerRoutes(routes, customer)

	// Address Route & Injection
	address := injection.AddressInjection(db)
	AddressRoutes(routes, address)

	return routes
}
