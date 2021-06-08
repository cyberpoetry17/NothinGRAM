package repository

import (
	"fmt"

	"github.com/cyberpoetry17/UserAPI/data"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type databaseRepositories struct {
	User2 IUserRepository

	database *gorm.DB
}

func SetRepositories(host, dbUser, dbName, password, dbPort string) (*databaseRepositories, error) {
	databaseUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, dbUser, dbName, password, dbPort)

	database, err := gorm.Open(postgres.Open(databaseUri), &gorm.Config{})

	if err != nil {
		panic("aaaa")
	} else {
		fmt.Printf("Successfully connected to your database GOPHER!!!")
	}

	return &databaseRepositories{
		User2:    NewUserRepository(database), //pravi repozitorijum za Usera2
		database: database,
	}, nil

}

func (repository *databaseRepositories) Automigrate() error {
	return repository.database.AutoMigrate(&data.User2{})
}
