package repository

import (
	entity "product/entity"

	"github.com/jinzhu/gorm"
)

type CategoryRepositoryContract interface {
	GetAll() []entity.Category
	GetByID(uint) entity.Category
	GetByName(string) entity.Category

	Create(entity.Category) (entity.Category, error)
	Update(entity.Category) (entity.Category, error)
	Delete(uint) error
}

type CategoryRepository struct {
	DB *gorm.DB
}

func ProviderCategoryRepository(DB *gorm.DB) CategoryRepository {
	return CategoryRepository{DB: DB}
}

func (c *CategoryRepository) GetAll() []entity.Category {

	var categories []entity.Category

	c.DB.Find(&categories)

	return categories
}

func (c *CategoryRepository) GetByID(id uint) entity.Category {

	var category entity.Category

	c.DB.Where("id=?", id).Find(&category)

	return category
}

func (c *CategoryRepository) GetByName(name string) entity.Category {

	var category entity.Category

	c.DB.Where("name=?", name).Find(&category)

	return category
}

func (c *CategoryRepository) Create(category entity.Category) (entity.Category, error) {

	err := c.DB.Create(&category).Error

	return category, err
}

func (c *CategoryRepository) Update(category entity.Category) (entity.Category, error) {

	err := c.DB.Model(&category).Where("id=?", category.ID).Update(&category).Error

	return category, err
}

func (c *CategoryRepository) Delete(id uint) error {

	var category entity.Category

	err := c.DB.Model(&category).Where("id=?", id).Updates(map[string]interface{}{
		"is_active": false,
	}).Error

	return err
}
