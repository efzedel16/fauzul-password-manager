package formatters

import "FauzulPasswordManager/entities"

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
	Address       string `json:"address"`
	Authorization string `json:"authorization"`
}

func UserFormat(data entities.User) UserFormatter {
	return UserFormatter{
		Id:       data.Id,
		FullName: data.FullName,
		Address:  data.Address,
		Email:    data.Email,
	}
}

func AllUsersFormat(datas []entities.User) []UserFormatter {
	var formatter []UserFormatter
	for _, data := range datas {
		formatter = append(formatter, UserFormat(data))
	}

	return formatter
}

func UserSignInFormat(data entities.User, token string) UserSignInFormatter {
	return UserSignInFormatter{
		Id:            data.Id,
		FullName:      data.FullName,
		Address:       data.Address,
		Email:         data.Email,
		Authorization: token,
	}
}
