package entity

import (
	"book-author/common/auth"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string `gorm:"not null;uniqueIndex" json:"username,omitempty" form:"username" valid:"required~Your username is required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email,omitempty" form:"email" valid:"required~Your email is required, email~Invalid email format,email~Invalid format email"`
	Password string `gorm:"not null" json:"password,omitempty" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(user)
	if errCreate != nil {
		err = errCreate
		return err
	}
	hash, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash
	return
}
