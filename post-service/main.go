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

func initializeRepository(database *gorm.DB) *repository.PostRepo {
	return &repository.PostRepo{Database: database}
}

func initializeServices(repo *repository.PostRepo) *services.PostService {
	return &services.PostService{Repo: repo}
}

func initializeHandlers(service *services.PostService) *handlers.PostHandler {
	return &handlers.PostHandler{Service: service}
}
func handleFunc(handler *handlers.PostHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handler.Hello).Methods("GET")
	router.HandleFunc("/", handler.CreatePost).Methods("POST")
	router.HandleFunc("/verify/{description}", handler.Verify).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORTP")), router))
}

func init() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {

	host, _ := os.LookupEnv("HOSTP")
	dbPort, _ := os.LookupEnv("DATABASE_PORTP")
	dbName, _ := os.LookupEnv("NAMEP")
	password, _ := os.LookupEnv("PASSWORDP")
	dbUser, _ := os.LookupEnv("USERP")
	//host, dbUser, dbName, password, dbPort string)
	db := repository.SetRepositoriesAndDatabase(host, dbUser, dbName, password, dbPort) //ovo je baza

	repo := initializeRepository(db)
	service := initializeServices(repo)
	handler := initializeHandlers(service)
	handleFunc(handler)
}
