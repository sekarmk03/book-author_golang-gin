package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Publisher struct {
	GormModel
	Name  string `gorm:"not null" json:"name" valid:"required~Name is required"`
	City  string `gorm:"not null" json:"city" valid:"required~City is required"`
	Books []Book `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"books"`
}

func (publisher *Publisher) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(publisher)
	if errCreate != nil {
		err = errCreate
		return err
	}
	return
}
