package model

import "github.com/jinzhu/gorm"

type Subscription struct {
	gorm.Model
	EventID    uint
	UserInfoID uint
}
