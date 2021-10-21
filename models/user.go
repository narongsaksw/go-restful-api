package models

import (
	"time"

	"github.com/matthewhartstonge/argon2"
	"gorm.io/gorm"
)

type User struct {
	ID			uint		`json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Fullname 	string		`json:"fullName" gorm:"type:varchar(255);not null"`
	Email		string		`json:"email" gorm:"type:varchar(255);not null;unique"`
	Password	string		`json:"-" gorm:"type:varchar(255);not null"`
	IsAdmin		bool		`json:"isAdmin" gorm:"type:bool;default:false"`
	CreatedAt   time.Time	`json:"createdAt"`
  	UpdatedAt   time.Time	`json:"updatedAt"`
}


func (user *User) BeforeCreate(db *gorm.DB) error {
	user.Password = hasPassword(user.Password)
	return nil
}

func hasPassword(password string) string {
	argon := argon2.DefaultConfig()
    encoded, _ := argon.HashEncoded([]byte(password))
	return string(encoded)
}