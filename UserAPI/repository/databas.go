package repository

import (
	"fmt"

	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
}

// type databaseRepositories struct {
// 	repo     UserRepo
// 	database *gorm.DB
// }

func SetRepositoriesAndDatabase(host, dbUser, dbName, password, dbPort string) *gorm.DB {
	databaseUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, dbUser, dbName, password, dbPort)

	database, err := gorm.Open(postgres.Open(databaseUri), &gorm.Config{})

	if err != nil {
		panic("aaaa gopher error!!")
	} else {
		fmt.Printf("Successfully connected to your database GOPHER!!!")
	}
	database.AutoMigrate(&data.User2{})

	users := []data.User2{
		//{Name: "Frodo", Surname: "Torbar", Email: "baggins@gmail.com", Username: "Saviour of the Middle Earth", Private: true, DateOfBirth: "12345", Gender: 1, PhoneNumber: "003345", Website: "OneRingToRuleThemAllButMe.com", Taggable: false, ReceiveNotifications: false, Password: "mypreci0us", Verified: true, Biography: "true", Role: 1},
	}
	for _, user := range users {
		database.Create(&user)
	}

	return database
}
