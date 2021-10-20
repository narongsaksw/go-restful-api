package models

import (
	"time"

	"github.com/matthewhartstonge/argon2"
	"gorm.io/gorm"
)

type User struct {
	ID			uint		`gorm:"primaryKey;autoIncrement;not null"`
	Fullname 	string		`gorm:"type:varchar(255);not null"`
	Email		string		`gorm:"type:varchar(255);not null;unique"`
	Password	string		`gorm:"type:varchar(255);not null"`
	IsAdmin		bool		`gorm:"type:bool;default:false"`
	CreatedAt   time.Time
  	UpdatedAt   time.Time
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