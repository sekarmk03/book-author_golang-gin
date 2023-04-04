package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Book struct {
	GormModel
	Isbn        string `gorm:"not null" json:"isbn" valid:"required~Isbn is required"`
	Title       string `gorm:"not null" json:"title" valid:"required~Title is required"`
	AuthorId    uint   `gorm:"not null" json:"author_id"`
	PublisherId uint   `gorm:"not null" json:"publisher_id"`
	Author      *Author
	Publisher   *Publisher
}

func (book *Book) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(book)
	if errCreate != nil {
		return errCreate
	}
	return
}
