package repository

import (
	entity "product/entity"

	"github.com/jinzhu/gorm"
)

type ProductRepositoryContract interface {
	GetAll() []entity.Product
	GetByID(uint) entity.Product
	GetByName(string) entity.Product
	GetByCategoryID(uint) []entity.Product

	Create(entity.Product) (entity.Product, error)
	Update(entity.Product) (entity.Product, error)
	Delete(uint) error
}

type ProductRepository struct {
	DB *gorm.DB
}

func ProviderProductRepository(DB *gorm.DB) ProductRepository {
	return ProductRepository{DB: DB}
}

func (p *ProductRepository) GetAll() []entity.Product {

	var products []entity.Product

	p.DB.Find(&products)

	return products
}

func (p *ProductRepository) GetByID(id uint) entity.Product {

	var product entity.Product

	p.DB.Where("id=?", id).Find(&product)

	return product
}

func (p *ProductRepository) GetByName(name string) entity.Product {

	var product entity.Product

	p.DB.Where("name=?", name).Find(&product)

	return product
}

func (p *ProductRepository) GetByCategoryID(category_id uint) []entity.Product {

	var products []entity.Product

	p.DB.Where("category_id=?", category_id).Find(&products)

	return products

}

func (p *ProductRepository) Create(product entity.Product) (entity.Product, error) {

	err := p.DB.Create(&product).Error

	return product, err
}

func (p *ProductRepository) Update(product entity.Product) (entity.Product, error) {

	err := p.DB.Model(&product).Where("id=?", product.ID).Update(&product).Error

	return product, err
}

func (p *ProductRepository) Delete(id uint) error {

	var product entity.Product

	err := p.DB.Model(&product).Where("id=?", id).Updates(map[string]interface{}{
		"is_active": false,
	}).Error

	return err
}
