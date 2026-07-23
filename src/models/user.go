package models

import (
	gorm "gorm.io/gorm"
)

type Gender uint8
type MaritalStatus uint8

const (
	MALE Gender = iota + 1
	FEMALE
	OTHER
)

const (
	SINGLE              MaritalStatus = iota + 1
	MARRIED            
	DIVORCED           
	WIDOWED            
	SEPARATED          
	DOMESTICPARTNERSHIP
)

var GenderResolutionMap = map[Gender]string{
	MALE: "male",
	FEMALE: "female",
	OTHER: "other",
}

var GenderChoices = map[string]Gender{
	"male": MALE,
	"female": FEMALE, 
	"other": OTHER,
}

var MaritalStatusChoices = map[string]MaritalStatus{
	"single": SINGLE,
	"married": MARRIED,
	"divorced": DIVORCED,
	"widowed": WIDOWED,
	"separated": SEPARATED,
	"partnership/common-Law": DOMESTICPARTNERSHIP,
}


var MaritalStatusResolutionMap = map[MaritalStatus]string{
	SINGLE: "single",
	MARRIED: "married",
	DIVORCED: "divorced",
	WIDOWED: "widowed",
	SEPARATED: "separated",
	DOMESTICPARTNERSHIP: "partnership/common-Law",
}

type User struct {
	gorm.Model
	FirstName    string        `gorm:"type:varchar(30);not null"`
	LastName     string        `gorm:"type:varchar(30);not null"`
	Gender       Gender           `gorm:"type:int;not null"`
	Email        string        `gorm:"type:varchar(254);not null;uniqueIndex:idx_unique_emails"`
	PasswordHash string        `gorm:"type:varchar(255);not null"`
	DateOfBirth  string     `gorm:"type:date;not null"`
	MaritalStatus MaritalStatus          `gorm:"type:int;not null"`
	
	RefreshToken *string       `gorm:"type:text"` 

	// Users who follow this user (Lookups read from FollowingID to discover FollowerID keys)
	Followers    []User        `gorm:"many2many:followers;foreignKey:ID;joinForeignKey:FollowingID;references:ID;joinReferences:FollowerID"`
	
	// Users this user is actively following (Lookups read from FollowerID to discover FollowingID keys)
	Following    []User        `gorm:"many2many:followers;foreignKey:ID;joinForeignKey:FollowerID;references:ID;joinReferences:FollowingID"` 

	OfficeDetails      []OfficeDetails      `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	ResidentialDetails []ResidentialDetails `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
}

/* Method to Decode Gender literal to respective String.
Returns "Unknown" on getting Invalid literal
*/
func (g Gender) String() string {
	value, exist := GenderResolutionMap[g]
	if !exist {
		return "Unknown"
	}
	
	return value
}


/* Method to Decode MaritalStatus literal to respective String.
Returns "Unknown" on getting Invalid literal
*/
func (m MaritalStatus) String() string {
	value, exist := MaritalStatusResolutionMap[m]
	if !exist {
		return "Unknown"
	}
	
	return value
}
