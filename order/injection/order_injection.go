package injection

import (
	"order/controller"
	repository "order/repository"
	"order/request"
	"order/service"

	"github.com/jinzhu/gorm"
)

func OrderInjection(db *gorm.DB) controller.OrderController {

	OrderRepository := repository.ProviderOrderRepository(db)
	ItemRepository := repository.ProviderItemRepository(db)

	CustomerRequest := request.ProviderCustomerRequest()
	ProductRequest := request.ProviderProductRequest()

	OrderService := service.ProviderOrderService(
		OrderRepository,
		ItemRepository,
		CustomerRequest,
		ProductRequest,
	)

	OrderController := controller.ProviderOrderController(OrderService)

	return OrderController
}
