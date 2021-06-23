package services

import (
	"FauzulPasswordManager/auth"
	"FauzulPasswordManager/entities"
	"FauzulPasswordManager/formatters"
	"FauzulPasswordManager/inputs"
	"FauzulPasswordManager/repositories"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	auth auth.AuthService
	repository repositories.UserRepository
}

func NewUserService(auth auth.AuthService, repository repositories.UserRepository) *service {
	return &service{auth, repository}
}

type UserService interface {
	SignUp(input inputs.SignUp) (formatters.UserFormatter, error)
	SignIn(input inputs.SignIn) (formatters.UserSignInFormatter, error)
	//Update(id int, input inputs.UpdateUser) (entities.User, error)
	//Delete(id int) (entities.User, error)
	GetAll() ([]formatters.UserFormatter, error)
}

func (s *service) SignUp(input inputs.SignUp) (formatters.UserFormatter, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	data := entities.User{
		FullName:     input.FullName,
		Address:      input.Address,
		Email:        input.Email,
		PasswordHash: string(passHash),
	}

	newData, err := s.repository.Create(data)
	if err != nil {
		return formatters.UserFormatter{}, err
	}

	formatter := formatters.UserFormat(newData)
	return formatter, nil
}

func (s *service) SignIn(input inputs.SignIn) (formatters.UserSignInFormatter, error) {
	data, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return formatters.UserSignInFormatter{}, err
	}

	if data.Id == 0 {
		return formatters.UserSignInFormatter{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(data.PasswordHash), []byte(input.Password)); err != nil {
		return formatters.UserSignInFormatter{}, err
	}

	token, _ := s.auth.Generate(data.Id)
	formatter := formatters.UserSignInFormat(data, token)
	return formatter, nil
}

//func (s *service) Update(id int, input inputs.UpdateUser) (entities.User, error) {
//	var dataUpdate = map[string]interface{}{}
//
//	if input.FullName != "" || len(input.FullName) == 0 {
//		dataUpdate["full_name"] = input.FullName
//	}
//
//	if input.Address != "" || len(input.Address) == 0 {
//		dataUpdate["address"] = input.Address
//	}
//
//	if input.Email != "" || len(input.Email) == 0 {
//		dataUpdate["email"] = input.Email
//	}
//
//	dataUpdate["updated_at"] = time.Now()
//}

//func (s *service) Delete(id int) (entities.User, error) {
//
//}


func (s *service) GetAll() ([]formatters.UserFormatter, error) {
	datas, err := s.repository.FindAll()
	if err != nil {
		return []formatters.UserFormatter{}, err
	}

	formatter := formatters.UsersFormat(datas)
	return formatter, nil
}