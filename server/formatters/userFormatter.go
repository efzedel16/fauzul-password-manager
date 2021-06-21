package formatters

import "PasswordManager/entities"

type UserFormatter struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

type UserSignInFormatter struct {
	Id            int    `json:"id"`
	FullName      string `json:"full_name"`
	Email         string `json:"email"`
	Authorization string `json:"authorization"`
}

func UserFormat(userData entities.User) UserFormatter {
	return UserFormatter{
		Id: userData.Id,
		FullName: userData.FullName,
		Address: userData.Address,
		Email: userData.Email,
	}
}

func AllUsersFormat(usersData []entities.User) []UserFormatter {
	var allUsersFormat []UserFormatter
	for _, user := range usersData {
		allUsersFormat = append(allUsersFormat, UserFormat(user))
	}

	return allUsersFormat
}