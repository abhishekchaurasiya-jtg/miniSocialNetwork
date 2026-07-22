package models

import "gorm.io/gorm"

type OfficeDetails struct {
	gorm.Model
	UserId uint					`gorm:"type:int;not null;index"`

	EmployeeId string			`gorm:"type:varchar(10);not null"`
	ContactNo string			`gorm:"type:varchar(20);not null"`
	Email string				`gorm:"type:varchar(254);not null"`
	Address string				`gorm:"type:varchar(500);not null"`
	City string					`gorm:"type:varchar(50);not null"`
	Country string				`gorm:"type:varchar(60);not null"`
	State string				`gorm:"type:varchar(50);not null"`
	Name string					`gorm:"type:varchar(255);not null"`
}
