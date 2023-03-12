package common

import (
	"log"

	"github.com/Ignosuke/not_a_game-api/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:******/mydb?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()
	log.Println("Migrating...")
	db.AutoMigrate(&models.Users{})
}
