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
	// database, err := gorm.Open(postgres.New(postgres.Config{
	// 	DSN:                  fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, dbUser, dbName, password, dbPort),
	// 	PreferSimpleProtocol: true,
	// }), &gorm.Config{})
	if err != nil {
		panic("aaaa gopher error!!")
	} else {
		fmt.Printf("Successfully connected to your database for posts!!!")
	}
	database.AutoMigrate(&data.Post{},&data.Tag{},&data.Comment{},&data.Like{},&data.Dislike{},&data.Media{},&data.Post{})
	return database
}
