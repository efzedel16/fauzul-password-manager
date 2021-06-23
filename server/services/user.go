package services

import (
	"FauzulPasswordManager/auth"
	"FauzulPasswordManager/entities"
	"FauzulPasswordManager/formatters"
	"FauzulPasswordManager/inputs"
	"FauzulPasswordManager/repositories"
	"errors"
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
	GetAll() ([]formatters.UserFormatter, error)
	GetById(id string) (formatters.UserFormatter, error)
}

func (s *userService) SignUp(input inputs.SignUp) (formatters.UserFormatter, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(input.Pass), bcrypt.DefaultCost)
	data := entities.User{
		FullName: input.FullName,
		Address:  input.Address,
		Email:    input.Email,
		PassHash: string(passHash),
		//CreatedAt: time.Now(),
		//UpdatedAt: time.Now(),
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

	if data.Id == 0 || len(data.FullName) <= 1 {
		return formatters.UserSignInFormatter{}, errors.New("wrong email / password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(data.PassHash), []byte(input.Pass)); err != nil {
		return formatters.UserSignInFormatter{}, errors.New("wrong email / password")
	}

	token, _ := s.authService.Generate(data.Id)
	formatter := formatters.UserSignInFormat(data, token)
	return formatter, nil
}

func (s *userService) GetAll() ([]formatters.UserFormatter, error) {
	datas, err := s.userRepository.FindAll()
	if err != nil {
		return []formatters.UserFormatter{}, err
	}

	formatter := formatters.AllUsersFormat(datas)
	return formatter, nil
}

func (s *userService) GetById(id string) (formatters.UserFormatter, error) {
	data, err := s.userRepository.FindById(id)
	if err != nil {
		return formatters.UserFormatter{}, err
	}

	//if data.Id == 0 {
	//	return formatters.UserFormatter{}, err
	//}

	formatter := formatters.UserFormat(data)
	return formatter, nil
}
