package injection

import (
	"product/controller"
	repository "product/repository"
	"product/service"

	"github.com/jinzhu/gorm"
)

func CategoryInjection(db *gorm.DB) controller.CategoryController {

	CategoryRepository := repository.ProviderCategoryRepository(db)
	ProductRepository := repository.ProviderProductRepository(db)

	CategoryService := service.ProviderCategoryService(CategoryRepository, ProductRepository)
	CategoryController := controller.ProviderCategoryController(CategoryService)

	return CategoryController
}
