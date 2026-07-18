package models

import "gorm.io/gorm"

type Followers struct {
	gorm.Model
	FollowerId   uint `gorm:"type:int;not null;index:idx_follower_history"`
	FolloweingId uint `gorm:"type:int;not null;index:idx_follower_history"`
}


