package models

import (
	"time"

	"gorm.io/gorm"
)

type Gender uint8
type MaritalStatus uint8

const (
	MALE Gender = iota + 1
	FEMALE
	OTHER
)

const (
	Single              MaritalStatus = iota + 1
	Married            
	Divorced           
	Widowed            
	Separated          
	DomesticPartnership
)

var genderResolutionMap = map[Gender]string{
	MALE: "male",
	FEMALE: "female",
	OTHER: "other",
}

var maritalStatusResolutionMap = map[MaritalStatus]string{
	Single: "single",
	Married: "married",
	Divorced: "divorced",
	Widowed: "widowed",
	Separated: "separated",
	DomesticPartnership: "partnership/common-Law",

}

type User struct {
	gorm.Model
	FirstName     string        `gorm:"type:varchar(30);not null"`
	LastName      string        `gorm:"type:varchar(30);not null"`
	Gender        Gender        `gorm:"type:int;not null"`
	Email         string        `gorm:"type:varchar(254);not null;unique"`
	PasswordHash  string        `gorm:"type:varchar(255);not null"`
	DateOfBirth   *time.Time    `gorm:"type:date;not null"`
	MaritalStatus MaritalStatus `gorm:"type:int;not null"`
	RefreshToken  string        `gorm:"type:varchar(255);not null"`

	// relation to office (one to many)
	// this automatically filters the record on preload where deletedAt is not null
OfficeDetails      []OfficeDetails      `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
ResidentialDetails []ResidentialDetails `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
}

// repository method..
// method to get current office Address
func (u *User) CurrentOffice() *OfficeAddress {
	if len(u.OfficeAddresses) > 0 {
		return &u.OfficeAddresses[0]
	}
	return nil
}

// Method of get Current Residential Address
func (u *User) CurrentResidentialAddress() *ResidentialAdrress {
	if len(u.ResidentialAddresses) > 0 {
		return &u.ResidentialAddresses[0]
	}
	return nil
}

// Method to Decode Gender literal to respective String.
// Returns "Unknown" on getting Invalid literal
func (g Gender) String() string {
	value, exist := genderResolutionMap[g]
	if !exist {
		return "Unknown"
	}
	
	return value
}


// Method to Decode MaritalStatus literal to respective String.
// Returns "Unknown" on getting Invalid literal
func (m MaritalStatus) String() string {
	value, exist := maritalStatusResolutionMap[m]
	if !exist {
		return "Unknown"
	}
	
	return value
}

