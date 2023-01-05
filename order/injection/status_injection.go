package injection

import (
	"order/controller"
	repository "order/repository"
	"order/service"

	"github.com/jinzhu/gorm"
)

func StatusInjection(db *gorm.DB) controller.StatusController {

	StatusRepository := repository.ProviderStatusRepository(db)
	OrderRepository := repository.ProviderOrderRepository(db)

	StatusService := service.ProviderStatusService(StatusRepository, OrderRepository)
	StatusController := controller.ProviderStatusController(StatusService)

	return StatusController
}
