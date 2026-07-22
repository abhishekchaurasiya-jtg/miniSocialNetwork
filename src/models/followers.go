package models

import (
	"errors"

	"gorm.io/gorm"
)

/*
using same name on index allows us to have composite index.
the composite unique index here ensures integrity + fast access for 1 -> M reation.
using extra index on second column fills the drawback and gives M <- 1 relation, very fast.
*/

type Follower struct {
	gorm.Model
	
	// The user who initiated the connection
	FollowerID  uint `gorm:"not null;uniqueIndex:idx_active_relation,where:deleted_at IS NULL"`
	Follower    User `gorm:"foreignKey:FollowerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	
	// The user who is being followed
	FollowingID uint `gorm:"not null;uniqueIndex:idx_active_relation,where:deleted_at IS NULL;index:idx_following_only"`
	Following   User `gorm:"foreignKey:FollowingID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}


func (f *Follower) BeforeCreate(tx *gorm.DB) (err error) {
	if f.FollowerID == f.FollowingID {
		return errors.New("users cannot follow themselves")
	}
	return nil
}