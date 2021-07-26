package repository

import (
	"fmt"

	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
}

func SetRepositoriesAndDatabase(host, dbUser, dbName, password, dbPort string) *gorm.DB {
	databaseUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, dbUser, dbName, password, dbPort)

	database, err := gorm.Open(postgres.Open(databaseUri), &gorm.Config{})
	// database, err := gorm.Open(postgres.New(postgres.Config{
	// 	DSN:                  fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, dbUser, dbName, password, dbPort),
	// 	PreferSimpleProtocol: true,
	// }), &gorm.Config{})
	if err != nil {
		panic("aaaa gopher error!!")
	} else {
		fmt.Printf("Successfully connected to your database GOPHER!!!")
	}
	database.AutoMigrate(&data.User2{})
	database.AutoMigrate(&data.Follower{})
	database.AutoMigrate(&data.Blocked{})
	database.AutoMigrate(&data.Muted{})
	database.AutoMigrate(&data.FollowerRequest{})
	database.AutoMigrate(&data.CloseFollower{})

	// users := []data.User2{
	// 	{Name: "Frodo", Surname: "Baggins", Email: "baggins@gmail.com", Username: "Saviour of the Middle Earth", Private: true, DateOfBirth: "12345", Gender: 1, PhoneNumber: "003345", Website: "OneRingToRuleThemAllButMe.com", Taggable: false, ReceiveNotifications: false, Password: "mypreci0us", Verified: true, Biography: "true", Role: 1},
	// 	//{Name: "Pera", Surname: "Peric", Email: "pera@gmail.com", Username: "Perica", Private: true, DateOfBirth: "12345", Gender: 1, PhoneNumber: "003345", Website: "pera.com", Taggable: false, ReceiveNotifications: false, Password: "pera123", Verified: true, Biography: "true", Role: 1},
	// }
	// for _, user := range users {

	// 	database.Create(&user)
	// }

	return database
}
