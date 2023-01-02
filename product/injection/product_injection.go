package injection

import (
	"product/controller"
	repository "product/repository"
	"product/service"

	"github.com/jinzhu/gorm"
)

func ProductInjection(db *gorm.DB) controller.ProductController {

	CategoryRepository := repository.ProviderCategoryRepository(db)
	ProductRepository := repository.ProviderProductRepository(db)

	ProductService := service.ProviderProductService(CategoryRepository, ProductRepository)
	ProductController := controller.ProviderProductController(ProductService)

	return ProductController
}
