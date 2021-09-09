package entity

import "time"

type Person struct {
	ID uint64 `json:"id" gorm:"primary_key;auto_increment"`
	FirstName string `json:"first_name" binding:"required" gorm:"type:varchar(32)"`
	LastName string `json:"last_name" binding:"required" gorm:"type:varchar(32)"`
	Age int8 `json:"age" binding:"gte=1,lte=130"`
	Email string `json:"email" binding:"required,email" gorm:"type:varchar(256)"`
}

type Video struct {
	ID uint64 `json:"id" gorm:"primary_key;auto_increment" `
	Title string `json:"title" binding:"min=2,max=10" validate:"is-cool" gorm:"type:varchar(100)"` // min 2 char, max 10 char
	Description string `json:"description" binding:"max=20" gorm:"type:varchar(200)"`
	URL string `json:"url" binding:"required,url" gorm:"type:varchar(256);UNIQUE"`
	Author Person `json:"author" binding:"required" gorm:"foreignkey:PersonID"`
	PersonID uint64 `json:"-"`
	CreatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
