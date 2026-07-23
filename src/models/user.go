package models

import (
	"strings"
	time "time"

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

var genderResolutionMap = map[Gender]string{
	MALE: "male",
	FEMALE: "female",
	OTHER: "other",
}

var maritalStatusResolutionMap = map[MaritalStatus]string{
	SINGLE: "single",
	MARRIED: "married",
	DIVORCED: "divorced",
	WIDOWED: "widowed",
	SEPARATED: "separated",
	DOMESTICPARTNERSHIP: "partnership/common-Law",

}

type User struct {
	gorm.Model
	FirstName    string        `gorm:"type:varchar(255);not null" validate:"required,gt=0,max=70"`
	LastName     *string       `gorm:"type:varchar(255)" validate:"omitempty,gt=0,max=100"`
	Gender       Gender        `gorm:"not null;check:chk_gender,gender BETWEEN	1 AND 3"`
	Email        string        `gorm:"type:varchar(255);not null;uniqueIndex:idx_unique_email_lower,expression:LOWER(email)"`
	PasswordHash string        `gorm:"type:varchar(255);not null"`
	DateOfBirth  time.Time     `gorm:"type:date;not null;check_date_birth <= (CURRENT_DATE - INTERVAL '12 years')"`		// 12 years constaraint
	MaritalStatus MaritalStatus`gorm:"not null;"`
	
	RefreshToken *string       `gorm:"type:text"`

	// Users who follow this user (Lookups read from FollowingID to discover FollowerID keys)
	Followers    []User        `gorm:"many2many:followers;foreignKey:ID;joinForeignKey:FollowingID;references:ID;joinReferences:FollowerID"`
	
	// Users this user is actively following (Lookups read from FollowerID to discover FollowingID keys)
	Following    []User        `gorm:"many2many:followers;foreignKey:ID;joinForeignKey:FollowerID;references:ID;joinReferences:FollowingID"` 

	OfficeDetails      []OfficeDetails      `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	ResidentialDetails []ResidentialDetails `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
}

// BeforeSave gorm hook
func (u *User)BeforeSave(tx *gorm.DB) (err error){
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	tx.Statement.SetColumn("Email", u.Email)
	return nil
}

/* Method to Decode Gender literal to respective String.
Returns "Unknown" on getting Invalid literal
*/
func (g Gender) String() string {
	value, exist := genderResolutionMap[g]
	if !exist {
		return "Unknown"
	}
	
	return value
}


/* Method to Decode MaritalStatus literal to respective String.
Returns "Unknown" on getting Invalid literal
*/
func (m MaritalStatus) String() string {
	value, exist := maritalStatusResolutionMap[m]
	if !exist {
		return "Unknown"
	}
	
	return value
}

