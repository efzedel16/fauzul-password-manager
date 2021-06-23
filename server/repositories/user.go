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
	//Update(id int, dataUpdate map[string]interface{}) (entities.User, error)
	//Delete(id int) (string, error)
	//FindAll() ([]entities.User, error)
	FindById(id int) (entities.User, error)
	FindByEmail(email string) (entities.User, error)
}

func (r *userRepository) Create(data entities.User) (entities.User, error) {
	if err := r.db.Create(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

//func (r *userRepository) Update(id int, dataUpdate map[string]interface{}) (entities.User, error) {
//	var data entities.User
//
//	if err := r.db.Where("id = ?", id).Error; err != nil {
//		return data, err
//	}
//
//	if err := r.db.Model(&data).Where("id = ?", id).Updates(&dataUpdate).Error; err != nil {
//		return data, err
//	}
//
//	return data, nil
//}

//func (r *userRepository) Delete(id int) (string, error) {
//	var data entities.User
//
//	if err := r.db.Where("id = ?", id).Delete(&data).Error; err != nil {
//		return "error", err
//	}
//
//	return "success", nil
//}

//func (r *userRepository) FindAll() ([]entities.User, error) {
//	var datas []entities.User
//
//	if err := r.db.Find(&datas).Error; err != nil {
//		return datas, err
//	}
//
//	return datas, nil
//}

func (r *userRepository) FindById(id int) (entities.User, error) {
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
