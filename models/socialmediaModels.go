package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name      string `gorm:"not null" valid:"required~your social media name is required"`
	SocialUrl string `gorm:"not null" valid:"required~your social media url is required"`
	UserID    uint
	User      *User
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
