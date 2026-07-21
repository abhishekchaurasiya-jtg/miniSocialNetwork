package dto

import (
	"app/src/models"
	"time"
)


type ResidentialDetailsRequest struct {
	Address    string `json:"address" binding:"required"`
	City       string `json:"city" binding:"required"`
	State      string `json:"state" binding:"required"`
	Country    string `json:"country" binding:"required"`
	ContactNo1 string `json:"contact_no_1" binding:"required"`
	ContactNo2 string `json:"contact_no_2" binding:"required"`
}

type OfficeDetailsRequest struct {
	EmployeeCode string `json:"employee_code" binding:"required"`
	Address      string `json:"address" binding:"required"`
	City         string `json:"city" binding:"required"`
	State        string `json:"state" binding:"required"`
	Country      string `json:"country" binding:"required"`
	ContactNo    string `json:"contact_no" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Name         string `json:"name" binding:"required"`
}

type UserDetailsRequest struct {
	FirstName      string                     `json:"first_name" binding:"required"`
	LastName       string                     `json:"last_name" binding:"required"`
	DateOfBirth    time.Time                  `json:"date_of_birth" binding:"required"`
	Gender         string                     `json:"gender" binding:"required"`
	MaritalStatus  string                     `json:"marital_status" binding:"required"`
	ResidentialDetails ResidentialDetailsRequest `json:"residential_details" binding:"required"`
	OfficeDetails      OfficeDetailsRequest      `json:"office_details" binding:"required"`
}

type CreateUserRequest struct {
	Email       string             `json:"email" binding:"required,email"`
	Password    string             `json:"password" binding:"required,min=8"`
	UserDetails UserDetailsRequest `json:"user_details" binding:"required"`
}

func (user_details *UserDetailsRequest) IsGenderValid() bool{
	_, exists := models.GenderChoices[user_details.Gender]
	return exists
}

func (user_details *UserDetailsRequest) GetEncodeGender() models.Gender{
	return models.GenderChoices[user_details.Gender]
}

func (user_details *UserDetailsRequest) IsMaritalStatusValid() bool {
	_, exists := models.MaritalStatusChoices[user_details.MaritalStatus]
	return exists
}

func (user_details *UserDetailsRequest) GetEncodeMaritalStatus() models.MaritalStatus {
	return models.MaritalStatusChoices[user_details.MaritalStatus]
}
// signup request schema finished.



// Login Request Schema
type LoginUserRequest struct {
	Email 		string				`json:"email" binding:"required,email"`
	Password 	string 				`json:"password" binding:"required,min=8"`
}