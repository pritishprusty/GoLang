package database

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/rahul10-pu/CIGO0122/models"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}
func Setup() {
	host := "localhost"
	port := "5432"
	dbName := "book"
	username := "postgres"
	passowrd := "postgres"
	args := "host=" + host + " port=" + port + " user=" + username + " dbname=" + dbName + " sslmode=disable password=" + passowrd
	db, err := gorm.Open("postgres", args)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(models.Book{}) // create table automatically in the db
	DB = db

}
