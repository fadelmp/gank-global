package repository

import (
	entity "order/entity"

	"github.com/jinzhu/gorm"
)

type StatusRepositoryContract interface {
	GetAll() []entity.Status
	GetByID(uint) entity.Status
	GetByName(string) entity.Status

	Create(entity.Status) (entity.Status, error)
	Update(entity.Status) (entity.Status, error)
	Delete(uint) error
}

type StatusRepository struct {
	DB *gorm.DB
}

func ProviderStatusRepository(DB *gorm.DB) StatusRepository {
	return StatusRepository{DB: DB}
}

func (s *StatusRepository) GetAll() []entity.Status {

	var statuses []entity.Status

	s.DB.Find(&statuses)

	return statuses
}

func (s *StatusRepository) GetByID(id uint) entity.Status {

	var status entity.Status

	s.DB.Where("id=?", id).Find(&status)

	return status
}

func (s *StatusRepository) GetByName(name string) entity.Status {

	var status entity.Status

	s.DB.Where("name=?", name).Find(&status)

	return status
}

func (s *StatusRepository) Create(status entity.Status) (entity.Status, error) {

	err := s.DB.Create(&status).Error

	return status, err
}

func (s *StatusRepository) Update(status entity.Status) (entity.Status, error) {

	err := s.DB.Model(&status).Where("id=?", status.ID).Update(&status).Error

	return status, err
}

func (c *StatusRepository) Delete(id uint) error {

	var status entity.Status

	err := c.DB.Model(&status).Where("id=?", id).Updates(map[string]interface{}{
		"is_active": false,
	}).Error

	return err
}
