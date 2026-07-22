package models

import "gorm.io/gorm"


type ResidentialAdrress struct {
	gorm.Model
	UserId uint				`gorm:"type:int;not null"`
	ContactNo1 string		`gorm:"type:varchar(20);not null"`
	ContactNo2 string		`gorm:"type:varchar(20)"`
	Address string			`gorm:"type:varchar(500);ot null"`
	City string				`gorm:"type:varchar(50);not null"`
	Country string			`gorm:"type:varchar(60);not null"`
	State string			`gorm:"type:varchar(50);not null"`
}

