package common

import (
	"log"

	"github.com/Ignosuke/not_a_game-api-rest/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:Hondacivic12_@/mydb?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()
	log.Println("Migranting...")
	db.AutoMigrate(&models.Users{})
}
