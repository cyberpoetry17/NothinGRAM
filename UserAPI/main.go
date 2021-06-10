package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/cyberpoetry17/NothinGRAM/UserAPI/handlers"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
)

func initializeRepository(database *gorm.DB) *repository.UserRepo {
	return &repository.UserRepo{Database: database}
}

func initializeServices(repo *repository.UserRepo) *services.UserService {
	return &services.UserService{Repo: repo}
}

func initializeHandlers(service *services.UserService) *handlers.UserHandler {
	return &handlers.UserHandler{Service: service}
}
func handleFunc(handler *handlers.UserHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handler.Hello).Methods("GET")
	router.HandleFunc("/", handler.CreateUser).Methods("POST")
	router.HandleFunc("/verify/{consumerId}", handler.Verify).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router))
}
func main() {

	host := os.Getenv("HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")
	dbUser := os.Getenv("USER")
	//host, dbUser, dbName, password, dbPort string)
	db := repository.SetRepositoriesAndDatabase(host, dbUser, dbName, password, dbPort) //ovo je baza

	repo := initializeRepository(db)
	service := initializeServices(repo)
	handler := initializeHandlers(service)
	handleFunc(handler)

}
