package main

import (
	"fmt"
	"os"

	"github.com/cyberpoetry17/UserAPI/data"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB
var err error

func main() {

	host := os.Getenv("HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")
	dbUser := os.Getenv("USER")

	databaseUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, dbUser, dbName, password, dbPort)

	//var dns string = "host=" + os.Getenv("HOST") + " " + "user=" + os.Getenv("USER") + " " + "password=" + os.Getenv("PASSWORD") + " " + "dbname=" + os.Getenv("NAME") + "port=" + os.Getenv("DATABASE_PORT")

	db, err := gorm.Open(postgres.Open(databaseUri), &gorm.Config{})

	if err != nil {
		panic("aaaa")
	} else {
		fmt.Printf("Successfully connected to your database GOPHER!!!")
	}

	db.AutoMigrate(data.User2{}) //baza mora znati kako da popuni sve ovo
	database = db

}
