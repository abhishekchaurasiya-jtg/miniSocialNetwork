package dto 

import (
	time "time"
)

// User Signup Response
type ResidentialDetailsResponse struct {
	Id         int    	`json:"id" validate:"required"`
	UserId     int    	`json:"user_id" validate:"required"`
	Address    string 	`json:"address" validate:"required"`
	City       string 	`json:"city" validate:"required"`
	State      string 	`json:"state" validate:"required"`
	Country    string 	`json:"country" validate:"required"`
	ContactNo1 string 	`json:"contact_no_1" validate:"required"`
	ContactNo2 string 	`json:"contact_no_2" validate:"required"`
}

type OfficeDetailsResponse struct {
	Id           int    `json:"id" validate:"required"`
	UserId       int    `json:"user_id" validate:"required"`
	EmployeeCode string `json:"employee_code" validate:"required"`
	Address      string `json:"address" validate:"required"`
	City         string `json:"city" validate:"required"`
	State        string `json:"state" validate:"required"`
	Country      string `json:"country" validate:"required"`
	ContactNo    string `json:"contact_no" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Name         string `json:"name" validate:"required"`
}

type TokenResponse struct {
	Token      string    `json:"key" validate:"required"`
	ExpiryTime time.Time `json:"expiry_time" validate:"required"`
}

type UserDetailsResponse struct {
	ID                 int                        `json:"id" validate:"required"`
	UserId             int                        `json:"user_id" validate:"required"`
	FirstName          string                     `json:"first_name" validate:"required"`
	LastName           string                     `json:"last_name" validate:"required"`
	DateOfBirth        string                     `json:"date_of_birth" validate:"required"`
	Gender             string                     `json:"gender" validate:"required"`
	MaritalStatus      string                     `json:"marital_status" validate:"required"`
	ResidentialDetails ResidentialDetailsResponse `json:"residential_details" validate:"required"` 
	OfficeDetails      OfficeDetailsResponse      `json:"office_details" validate:"required"`      
}


type UserCreateResponse struct {
	ID           int                 `json:"id" validate:"required"`
	UserId       int                 `json:"user_id" validate:"required"`
	Email        string              `json:"email" validate:"required,email,max=254"` 
	LastModified time.Time           `json:"last_modified" validate:"required"`
	UserDetail   UserDetailsResponse `json:"user_details" validate:"required"`  
	Token        TokenResponse       `json:"token" validate:"required"`          
}

// User Login response

type UserDetailsLoginResponse struct {
	ID 			 int				 `json:"id" validate:"required"`
	Email 		 string				 `json:"email" validate:"required"`
	FirstName    string				 `json:"first_name" validate:"required"`
	LastName     string				 `json:"last_name" validate:"required"`
	LastModified time.Time  		 `json:"last_modified" validate:"required"`
}

type UserLoginResponse struct {
	Token TokenResponse				 `json:"token" validate:"required"`
	User  UserDetailsLoginResponse   `json:"user" validate:"required"`
}


type UpdateUserResponse struct {
	ID           int                 `json:"id" validate:"required"`
	UserId       int                 `json:"user_id" validate:"required"`
	Email        string              `json:"email" validate:"required,email,max=254"` 
	LastModified time.Time           `json:"last_modified" validate:"required"`
	UserDetail   UserDetailsResponse `json:"user_details" validate:"required"`  
}

type DeleteUserResponse struct {
	ID           int                 `json:"id" validate:"required"`
	UserId       int                 `json:"user_id" validate:"required"`
	Email        string              `json:"email" validate:"required,email,max=254"` 
	LastModified time.Time           `json:"last_modified" validate:"required"`
	UserDetail   UserDetailsResponse `json:"user_details" validate:"required"`  
}

type GetUserResponse struct {
	ID           int                 `json:"id" validate:"required"`
	UserId       int                 `json:"user_id" validate:"required"`
	Email        string              `json:"email" validate:"required,email,max=254"` 
	LastModified time.Time           `json:"last_modified" validate:"required"`
	UserDetail   UserDetailsResponse `json:"user_details" validate:"required"`  
}

type GetActiveUsersResponseItem struct {
	UserID int `json:"userId" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}
