package models

import (
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Uname string `json:"uname"`
	Email string `json:"email"`
}
