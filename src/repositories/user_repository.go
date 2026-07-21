package repositories

import (
	"fmt"

	gorm "gorm.io/gorm"

	models "app/src/models"
)

type UserRepository interface {
	AppendSingleNewUser(user *models.User) error
	GetUsersCountByEmail(email string) (int64, error)
	UpdateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (user_repo *userRepository)AppendSingleNewUser(user *models.User) error {
	result := user_repo.db.Create(user)
	if err := result.Error; err != nil {
		return err
	}else if result.RowsAffected < 1 {
		return fmt.Errorf("Failed to Insert New Record.")
	}
	return nil
}

func (user_repo *userRepository)GetUsersCountByEmail(email string) (int64, error) {
	var user_count int64;
	if err := user_repo.db.Model(&models.User{}).Where("email = ?", email).Count(&user_count).Error; err != nil {
		return 0, err
	}
	return user_count, nil
}

func (user_repo *userRepository)UpdateUser(user *models.User) error{
	if err := user_repo.db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (user_repo *userRepository)GetUserByEmail(email string) (*models.User, error) {
	var user *models.User
	if err := user_repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}