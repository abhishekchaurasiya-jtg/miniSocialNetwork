package repositories

import (
	"fmt"

	gorm "gorm.io/gorm"

	models "app/src/models"
)

type AuthRepository interface {
	AppendSingleNewUser(user *models.User) error
	GetUsersCountByEmail(email string) (int64, error)
	UpdateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	PreloadAddresses(user *models.User) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (auth_repo *authRepository)AppendSingleNewUser(user *models.User) error {
	result := auth_repo.db.Create(user)
	if err := result.Error; err != nil {
		return err
	}else if result.RowsAffected < 1 {
		return fmt.Errorf("Failed to Insert New Record.")
	}
	return nil
}

func (auth_repo *authRepository)GetUsersCountByEmail(email string) (int64, error) {
	var user_count int64;
	if err := auth_repo.db.Model(&models.User{}).Where("email = ?", email).Count(&user_count).Error; err != nil {
		return 0, err
	}
	return user_count, nil
}

func (auth_repo *authRepository)UpdateUser(user *models.User) error{
	if err := auth_repo.db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (auth_repo *authRepository)GetUserByEmail(email string) (*models.User, error) {
	var user *models.User
	if err := auth_repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (auth_repo *authRepository)PreloadAddresses(user *models.User) error {
	return auth_repo.db.Preload("OfficeDetails").Preload("ResidentialDetails").First(user, user.ID).Error
}
