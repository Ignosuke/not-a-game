package models

import (
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	ID    int64  `json:"id" gorm:"primary_key;auto_increment"`
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Uname string `json:"uname"`
	Email string `json:"email"`
}
