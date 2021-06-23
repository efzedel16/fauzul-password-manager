package services

import (
	"FauzulPasswordManager/auth"
	"FauzulPasswordManager/entities"
	"FauzulPasswordManager/formatters"
	"FauzulPasswordManager/inputs"
	"FauzulPasswordManager/repositories"
	"fmt"
	"time"
)

type passService struct {
	authService    auth.AuthService
	passRepository repositories.PassRepository
}

func NewPassService(authService auth.AuthService, passRepository repositories.PassRepository) *passService {
	return &passService{authService, passRepository}
}

type PassService interface {
	Create(userId int, input inputs.CreatePass) (formatters.PassFormatter, error)
	Update(id int, input inputs.UpdatePass) (formatters.PassFormatter, error)
	Delete(id int) (interface{}, error)
	GetAllByUserId(userId int) ([]formatters.PassFormatter, error)
	GetById(id int) (formatters.PassFormatter, error)
}

func (s *passService) Create(userId int, input inputs.CreatePass) (formatters.PassFormatter, error) {
	data := entities.Password{
		UserId: userId,
		Web:    input.Web,
		Pass:   input.Pass,
	}

	newData, err := s.passRepository.Create(data)
	if err != nil {
		return formatters.PassFormatter{}, err
	}

	formatter := formatters.PassFormat(newData)
	return formatter, nil
}

func (s *passService) Update(id int, input inputs.UpdatePass) (formatters.PassFormatter, error) {
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

func (s *passService) Delete(id int) (interface{}, error) {
	data, err := s.passRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	if data.Id == 0 {
		return nil, err
	}

	pass, err := s.passRepository.Delete(id)
	if err != nil {
		return nil, err
	}

	if pass == "error" {
		return pass, err
	}

	msg := fmt.Sprintf("successfully delete password id %v", id)
	return msg, nil
}

func (s *passService) GetAllByUserId(userId int) ([]formatters.PassFormatter, error) {
	datas, err := s.passRepository.FindByUserId(userId)
	if err != nil {
		return []formatters.PassFormatter{}, err
	}

	formatter := formatters.AllPasssFormat(datas)
	return formatter, nil
}

func (s *passService) GetById(id int) (formatters.PassFormatter, error) {
	data, err := s.passRepository.FindById(id)
	if err != nil {
		return formatters.PassFormatter{}, err
	}

	if data.Id != 0 {
		return formatters.PassFormatter{}, err
	}

	return formatters.PassFormat(data), nil
}
