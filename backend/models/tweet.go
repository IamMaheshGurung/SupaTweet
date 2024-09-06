package models

import "github.com/jinzhu/gorm"

type Tweet struct {
	gorm.Model
	Content string `json:"Content"`
	UserId  uint   `json:"user_id"`
}
