package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	Message     string `gorm:"not null" json:"message" form:"message"`
	UserID      uint
	Photo		*Photo
	SocialMedia *SocialMedia
}

func (com *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(com)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
