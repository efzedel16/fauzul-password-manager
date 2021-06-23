package repositories

import (
	"FauzulPasswordManager/entities"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

type UserRepository interface {
	Create(data entities.User) (entities.User, error)
	FindAll() ([]entities.User, error)
	FindById(id string) (entities.User, error)
	FindByEmail(email string) (entities.User, error)
}

func (r *userRepository) Create(data entities.User) (entities.User, error) {
	if err := r.db.Create(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (r *userRepository) FindAll() ([]entities.User, error) {
	var datas []entities.User

	if err := r.db.Find(&datas).Error; err != nil {
		return datas, err
	}

	return datas, nil
}

func (r *userRepository) FindById(id string) (entities.User, error) {
	var data entities.User

	if err := r.db.Where("id = ?", id).Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (r *userRepository) FindByEmail(email string) (entities.User, error) {
	var data entities.User

	if err := r.db.Where("email = ?", email).Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}
