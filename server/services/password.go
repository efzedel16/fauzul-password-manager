package services

import (
	"FauzulPasswordManager/entities"
	"FauzulPasswordManager/formatters"
	"FauzulPasswordManager/inputs"
	"FauzulPasswordManager/repositories"
	"fmt"
	"time"
)

type passService struct {
	passRepository repositories.PassRepository
}

func NewPassService(passRepository repositories.PassRepository) *passService {
	return &passService{passRepository}
}

type PassService interface {
	Create(userId int, input inputs.CreatePass) (formatters.PassFormatter, error)
	Update(id string, input inputs.UpdatePass) (formatters.PassFormatter, error)
	Delete(id string) (string, error)
	GetAllByUserId(userId string) ([]formatters.PassFormatter, error)
	GetById(id string) (formatters.PassFormatter, error)
}

func (s *passService) Create(userId int, input inputs.CreatePass) (formatters.PassFormatter, error) {
	data := entities.Password{
		UserId:    userId,
		Web:       input.Web,
		Pass:      input.Pass,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	newData, err := s.passRepository.Create(data)
	if err != nil {
		return formatters.PassFormatter{}, err
	}

	formatter := formatters.PassFormat(newData)
	return formatter, nil
}

func (s *passService) Update(id string, input inputs.UpdatePass) (formatters.PassFormatter, error) {
	var dataUpdate = map[string]interface{}{}
	data, err := s.passRepository.FindById(id)
	if err != nil {
		return formatters.PassFormatter{}, err
	}

	if data.Id == 0 {
		return formatters.PassFormatter{}, err
	}

	if input.Web != "" || len(input.Web) == 0 {
		dataUpdate["web"] = input.Web
	}

	if input.Pass != "" || len(input.Pass) == 0 {
		dataUpdate["pass"] = input.Pass
	}

	dataUpdate["updated_at"] = time.Now()
	pass, err := s.passRepository.Update(id, dataUpdate)
	if err != nil {
		return formatters.PassFormatter{}, err
	}

	formatter := formatters.PassFormat(pass)
	return formatter, nil
}

func (s *passService) Delete(id string) (string, error) {
	pass, err := s.passRepository.Delete(id)
	if err != nil || pass == "error" {
		return pass, err
	}

	msg := fmt.Sprintf("successfully delete password id %v", id)
	return msg, nil
}

func (s *passService) GetAllByUserId(userId string) ([]formatters.PassFormatter, error) {
	datas, err := s.passRepository.FindByUserId(userId)
	if err != nil {
		return []formatters.PassFormatter{}, err
	}

	formatter := formatters.AllPasssFormat(datas)
	return formatter, nil
}

func (s *passService) GetById(id string) (formatters.PassFormatter, error) {
	data, err := s.passRepository.FindById(id)
	if err != nil {
		return formatters.PassFormatter{}, err
	}

	if data.Id == 0 {
		return formatters.PassFormatter{}, err
	}

	formatter := formatters.PassFormat(data)
	return formatter, nil
}
