package service

import (
	"errors"
	"order/config"
	"order/dto"
	entity "order/entity"
	"order/mapper"
	repository "order/repository"
)

type StatusServiceContract interface {
	GetAll() []dto.Status
	GetByID(id uint) dto.Status

	Create(entity.Status) (dto.Status, error)
	Update(entity.Status) (dto.Status, error)
	Delete(id uint) error
}

type StatusService struct {
	StatusRepository repository.StatusRepository
	OrderRepository  repository.OrderRepository
}

func ProviderStatusService(
	s repository.StatusRepository,
	o repository.OrderRepository,
) StatusService {
	return StatusService{
		StatusRepository: s,
		OrderRepository:  o,
	}
}

func (s *StatusService) GetAll() []dto.Status {

	statuses := s.StatusRepository.GetAll()

	return mapper.ToStatusDtoList(statuses)
}

func (s *StatusService) GetByID(id uint) dto.Status {

	status := s.StatusRepository.GetByID(id)

	return mapper.ToStatusDto(status)
}

func (s *StatusService) Create(dto dto.Status) (dto.Status, error) {

	if !s.CheckName(dto) {
		return dto, errors.New(config.StatusExists)
	}

	status_entity := mapper.ToStatusEntity(dto)

	status, err := s.StatusRepository.Create(status_entity)

	return mapper.ToStatusDto(status), err
}

func (s *StatusService) Update(dto dto.Status) (dto.Status, error) {

	if !s.CheckID(dto.ID) {
		return dto, errors.New(config.StatusNotFound)
	}

	status_entity := mapper.ToStatusEntity(dto)

	status, err := s.StatusRepository.Update(status_entity)

	return mapper.ToStatusDto(status), err
}

func (s *StatusService) Delete(id uint) error {

	if !s.CheckID(id) {
		return errors.New(config.StatusNotFound)
	}

	if !s.CheckOrder(id) {
		return errors.New(config.DeleteStatusFailed)
	}

	return s.StatusRepository.Delete(id)
}

func (s *StatusService) CheckName(dto dto.Status) bool {

	status_name := dto.Name

	status_data := s.StatusRepository.GetByName(status_name)

	if status_data.ID != 0 && status_data.IsActive {
		return false
	}

	return true
}

func (s *StatusService) CheckID(id uint) bool {

	status_data := s.StatusRepository.GetByID(id)

	if status_data.ID == 0 || !status_data.IsActive {
		return false
	}

	return true
}

func (s *StatusService) CheckOrder(id uint) bool {

	status_data := s.OrderRepository.GetByStatusID(id)

	if len(status_data) > 0 {
		return false
	}

	return true
}
