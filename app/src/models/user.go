package models

import (
	"fmt"
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
	Single              MaritalStatus = iota
	Married             MaritalStatus = iota
	Divorced            MaritalStatus = iota
	Widowed             MaritalStatus = iota
	Separated           MaritalStatus = iota
	DomesticPartnership MaritalStatus = iota
)

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
	OfficeAddresses      []OfficeAddress      `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	ResidentialAddresses []ResidentialAdrress `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
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

// Method to convert the integer literal to respective fullform for logging perpose
func (g Gender) ToString() (string, error) {
	switch g {
	case MALE:
		return "Male", nil
	case FEMALE:
		return "female", nil
	case OTHER:
		return "other", nil
	default:
		return "", fmt.Errorf("Invalid Literal for Gender")
	}
}

// method to conver the integer martial status back to string for logging perpose
func (m MaritalStatus) ToString() (string, error) {
	switch m {
	case Single:
		return "Single", nil
	case Married:
		return "Married", nil
	case Divorced:
		return "Divorced", nil
	case Widowed:
		return "Widowed", nil
	case Separated:
		return "Separated", nil
	case DomesticPartnership:
		return "Registered Partnership/Common-Law", nil
	default:
		return "", fmt.Errorf("invalid literal for MaritalStatus: %d", m)
	}
}
