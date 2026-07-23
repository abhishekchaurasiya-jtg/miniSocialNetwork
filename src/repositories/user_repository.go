package repositories

import (
	"log"

	gorm "gorm.io/gorm"

	dto "app/src/dto"
	models "app/src/models"
	responses "app/src/responses"
)


type UserRepository interface {
	DeleteUser(user *models.User) error
	GetActiveUsers(pageNo, pageSize int) (*[]dto.GetActiveUsersResponseItem, error)
	Save(user *models.User) error
	GetUserByID(userId int) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository)DeleteUser(user *models.User) error {
	result := ur.db.Delete(&user)
	if err := result.Error; err != nil {
		return err
	}
	if result.RowsAffected != 1 {
		return responses.DBIneffectiveOperation
	}
	return nil
}

func (ur *userRepository)GetActiveUsers(pageNo int, pageSize int) (*[]dto.GetActiveUsersResponseItem, error) {
	var results []dto.GetActiveUsersResponseItem

	offSet := (pageNo - 1) * pageSize
	err := ur.db.Model(&models.User{}).Order("id ASC").Limit(pageNo).Offset(offSet).Scan(&results).Error
	if err != nil {
		return nil, responses.DBFailedToFetchRecords
	}
	return &results, nil
}

func (ur *userRepository)Save(user *models.User) error {
	if err := ur.db.Save(&user).Error; err != nil {
		return responses.DBFailedToUpdateRecord
	}
	return nil
}

func (ur *userRepository)GetUserByID(userID int) (*models.User, error) {
	var user *models.User
	err := ur.db.
		Preload("OfficeDetails").
        Preload("ResidentialDetails").
		Where("id = ?", userID).
		First(&user).Error
	if err != nil {
		return nil, responses.DBFailedToFetchRecords
	}
	log.Printf("recived values from userbyid %v", user)
	return user, nil
}