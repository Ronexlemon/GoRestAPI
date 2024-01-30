package model

import (
	"example/restapi/database"

	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	Content string `gorm:"type:text" json:"content"`
	UserID  uint
}

func (entry *Entry) Save() (*Entry, error) {
	err := database.Database.Create(&entry).Error
	if err != nil {
		return entry, nil
	}

	return entry, nil
}
