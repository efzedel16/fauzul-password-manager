package services

import (
	"FauzulPasswordManager/auth"
	"FauzulPasswordManager/entities"
	"FauzulPasswordManager/formatters"
	"FauzulPasswordManager/inputs"
	"FauzulPasswordManager/repositories"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	authService    auth.AuthService
	userRepository repositories.UserRepository
}

func NewUserService(authService auth.AuthService, userRepository repositories.UserRepository) *userService {
	return &userService{authService, userRepository}
}

type UserService interface {
	SignUp(input inputs.SignUp) (formatters.UserFormatter, error)
	SignIn(input inputs.SignIn) (formatters.UserSignInFormatter, error)
	//Update(id int, input inputs.UpdateUser) (formatters.UserFormatter, error)
	//Delete(id int) (string, error)
	//GetAll() ([]formatters.UserFormatter, error)
	GetById(id int) (formatters.UserFormatter, error)
}

func (s *userService) SignUp(input inputs.SignUp) (formatters.UserFormatter, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(input.Pass), bcrypt.DefaultCost)
	data := entities.User{
		FullName: input.FullName,
		Address:  input.Address,
		Email:    input.Email,
		PassHash: string(passHash),
	}

	newData, err := s.userRepository.Create(data)
	if err != nil {
		return formatters.UserFormatter{}, err
	}

	formatter := formatters.UserFormat(newData)
	return formatter, nil
}

func (s *userService) SignIn(input inputs.SignIn) (formatters.UserSignInFormatter, error) {
	data, err := s.userRepository.FindByEmail(input.Email)
	if err != nil {
		return formatters.UserSignInFormatter{}, err
	}

	if data.Id == 0 {
		return formatters.UserSignInFormatter{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(data.PassHash), []byte(input.Pass)); err != nil {
		return formatters.UserSignInFormatter{}, err
	}

	token, _ := s.authService.Generate(data.Id)
	formatter := formatters.UserSignInFormat(data, token)
	return formatter, nil
}

//func (s *userService) Update(id int, input inputs.UpdateUser) (formatters.UserFormatter, error) {
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
//	data, err := s.userRepository.Update(id, dataUpdate)
//	if err != nil {
//		return formatters.UserFormatter{}, err
//	}
//
//	formatter := formatters.UserFormat(data)
//	return formatter, nil
//}

//func (s *userService) Delete(id int) (string, error) {
//	data, err := s.userRepository.Delete(id)
//	if err != nil || data == "error" {
//		return data, err
//	}
//
//	msg := fmt.Sprintf("successfully delete user id %v", id)
//	return msg, nil
//}

//func (s *userService) GetAll() ([]formatters.UserFormatter, error) {
//	datas, err := s.userRepository.FindAll()
//	if err != nil {
//		return []formatters.UserFormatter{}, err
//	}
//
//	formatter := formatters.AllUsersFormat(datas)
//	return formatter, nil
//}

func (s *userService) GetById(id int) (formatters.UserFormatter, error) {
	data, err := s.userRepository.FindById(id)
	if err != nil {
		return formatters.UserFormatter{}, err
	}

	if data.Id != 0 {
		return formatters.UserFormatter{}, err
	}

	return formatters.UserFormat(data), nil
}
