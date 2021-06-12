package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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
	router.HandleFunc("/register", handler.CreateUser).Methods("POST")
	router.HandleFunc("/verify/{userId}", handler.Verify).Methods("GET")
	router.HandleFunc("/login", handler.LoginUser).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router))
}

func init() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {

	host, _ := os.LookupEnv("HOST")
	dbPort, _ := os.LookupEnv("DATABASE_PORT")
	dbName, _ := os.LookupEnv("NAME")
	password, _ := os.LookupEnv("PASSWORD")
	dbUser, _ := os.LookupEnv("USER")
	//host, dbUser, dbName, password, dbPort string)
	db := repository.SetRepositoriesAndDatabase(host, dbUser, dbName, password, dbPort) //ovo je baza

	repo := initializeRepository(db)
	service := initializeServices(repo)
	handler := initializeHandlers(service)
	handleFunc(handler)
}
