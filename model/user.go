package model

import (
	"example/restapi/database"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"-"`
	Entries []Entry
}

func(user *User) Save() (*User,error){
	err:= database.Database.Create(&user).Error
	if err !=nil{
		return user, nil
	}

	return user,nil
}

func (user *User) BeforeSave(*gorm.DB)error{
	passwordHash,err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)

	if err !=nil{
		return nil
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}