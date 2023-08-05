package repository

import (
	"project-for-portfolioDEV/GOJWT-Auth/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	SaveUser(users model.Users)
	UpdateUser(users model.Users)
	DeleteUser(userId int) error
	FindUserById(users int) (model.Users, error)
	FindUserAll() []model.Users
	FindUserByUsername(username string) (model.Users, error)
}

type userRepository struct {
	Db *gorm.DB
}

func NewUserRepository(Db *gorm.DB) IUserRepository {
	return &userRepository{Db: Db}
}

// DeleteUser implements IUserRepository.
func (repo *userRepository) DeleteUser(userId int) error {
	var user model.Users

	result := repo.Db.Where("id = ?", userId).Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// FindUserAll implements IUserRepository.
func (repo *userRepository) FindUserAll() []model.Users {
	users := make([]model.Users, 0)

	result := repo.Db.Find(&users)
	_ = result

	return users
}

// FindUserById implements IUserRepository.
func (repo *userRepository) FindUserById(usersId int) (model.Users, error) {
	var user model.Users
	result := repo.Db.Find(&user, usersId)
	if result == nil {
		return user, result.Error
	}

	return user, nil
}

// FindUserByUsername implements IUserRepository.
func (repo *userRepository) FindUserByUsername(username string) (model.Users, error) {
	var user model.Users

	result := repo.Db.First(&user, "username = ?", username)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

// SaveUser implements IUserRepository.
func (repo *userRepository) SaveUser(users model.Users) {
	result := repo.Db.Create(&users)
	_ = result
}

// UpdateUser implements IUserRepository.
func (repo *userRepository) UpdateUser(users model.Users) {
	var updateUser = model.UpdateUserRequest{
		Id: users.Id,
		UserStandard: model.UserRequest{
			Username: users.Username,
			Email:    users.Email,
			Password: users.Password,
		},
	}
	result := repo.Db.Model(&users).Updates(updateUser)
	_ = result

}
