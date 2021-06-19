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

func initializeRepository(database *gorm.DB) (*repository.UserRepo, *repository.BlockedRepo) {
	return &repository.UserRepo{Database: database}, &repository.BlockedRepo{Database: database}
}

func initializeServices(repo *repository.UserRepo, repoBlocked *repository.BlockedRepo) (*services.UserService, *services.BlockedService) {
	return &services.UserService{Repo: repo}, &services.BlockedService{Repo: repoBlocked}
}

func initializeHandlers(service *services.UserService, serviceBlocked *services.BlockedService) (*handlers.UserHandler, *handlers.BlockedHandler) {
	return &handlers.UserHandler{Service: service}, &handlers.BlockedHandler{Service: serviceBlocked}
}
func handleFuncUser(handler *handlers.UserHandler, handlerBlocked *handlers.BlockedHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handler.Hello).Methods("GET")
	router.HandleFunc("/user/{userId}", handler.GetById).Methods("GET")
	router.HandleFunc("/register", handler.CreateUser).Methods("POST")
	router.HandleFunc("/update", handler.UpdateUser).Methods("POST")
	router.HandleFunc("/verify/{userId}", handler.Verify).Methods("GET")
	router.HandleFunc("/login", handler.LoginUser).Methods("POST")

	router.HandleFunc("/block", handlerBlocked.BlockUser).Methods("POST")
	router.HandleFunc("/unblock", handlerBlocked.UnblockUser).Methods("POST")
	router.HandleFunc("/allblockedusers/{userID}", handlerBlocked.GetAllBlockedUsers).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("USER_SERVICE_PORT")), router))
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

	repositoryUser, repositoryBlocked := initializeRepository(db)
	serviceUser, serviceBLocked := initializeServices(repositoryUser, repositoryBlocked)
	handlerUser, handlerBlocked := initializeHandlers(serviceUser, serviceBLocked)
	handleFuncUser(handlerUser, handlerBlocked)
}
