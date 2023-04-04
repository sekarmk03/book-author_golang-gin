package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Author struct {
	GormModel
	Name  string `gorm:"not null" json:"name" valid:"required~Name is required"`
	Email string `gorm:"not null" json:"email" valid:"required~Email is required"`
	Books []Book `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"books"`
}

func (author *Author) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(author)
	if errCreate != nil {
		err = errCreate
		return err
	}
	return
}
