package repository

import (
	entity "order/entity"

	"github.com/jinzhu/gorm"
)

type ItemRepositoryContract interface {
	GetAll() []entity.Item
	GetByOrderID(uint) []entity.Item
	GetByProductID(uint) []entity.Item

	Create(entity.Item) (entity.Item, error)
	Update(entity.Item) (entity.Item, error)
	Delete(uint) error
}

type ItemRepository struct {
	DB *gorm.DB
}

func ProviderItemRepository(DB *gorm.DB) ItemRepository {
	return ItemRepository{DB: DB}
}

func (i *ItemRepository) GetAll() []entity.Item {

	var items []entity.Item

	i.DB.Find(&items)

	return items
}

func (i *ItemRepository) GetByOrderID(order_id uint) []entity.Item {

	var items []entity.Item

	i.DB.Where("order_id=?", order_id).Find(&items)

	return items
}

func (i *ItemRepository) GetByProductID(product_id uint) []entity.Item {

	var items []entity.Item

	i.DB.Where("product_id=?", product_id).Find(&items)

	return items
}

func (i *ItemRepository) Create(item entity.Item) (entity.Item, error) {

	err := i.DB.Create(&item).Error

	return item, err
}

func (i *ItemRepository) Update(item entity.Item) (entity.Item, error) {

	err := i.DB.Model(&item).Where("id=?", item.ID).Update(&item).Error

	return item, err
}

func (i *ItemRepository) Delete(id uint) error {

	var item entity.Item

	err := i.DB.Model(&item).Where("id=?", id).Updates(map[string]interface{}{
		"is_active": false,
	}).Error

	return err
}
