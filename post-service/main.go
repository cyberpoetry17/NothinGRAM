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

func initializeRepository(database *gorm.DB) (*repository.PostRepo,*repository.TagRepo,*repository.CommentRepo) {
	return &repository.PostRepo{Database: database}, &repository.TagRepo{Database: database}, &repository.CommentRepo{Database: database}
}

func initializeServices(repoPost *repository.PostRepo, repoTag *repository.TagRepo, repoComment *repository.CommentRepo) (*services.PostService,*services.TagService,*services.CommentService) {
	return &services.PostService{Repo: repoPost}, &services.TagService{Repo: repoTag}, &services.CommentService{Repo: repoComment}
}

func initializeHandlers(servicePost *services.PostService,serviceTag *services.TagService, serviceComment *services.CommentService) (*handlers.PostHandler,*handlers.TagHandler,*handlers.CommentHandler) {
	return &handlers.PostHandler{Service: servicePost}, &handlers.TagHandler{Service: serviceTag}, &handlers.CommentHandler{Service: serviceComment}
}
func handleFunc(handler *handlers.PostHandler,tagHandler *handlers.TagHandler, commentHandler *handlers.CommentHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handler.Hello).Methods("GET")
	router.HandleFunc("/", handler.CreatePost).Methods("POST")
	router.HandleFunc("/verify/{description}", handler.Verify).Methods("GET")
	router.HandleFunc("/addTag/",tagHandler.CreateTag).Methods("POST")
	router.HandleFunc("/editTag/",tagHandler.EditTag).Methods("POST")
	router.HandleFunc("/removeTag/",tagHandler.DeleteTag).Methods("DELETE")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("POST_SERVICE_PORT")), router))
}

func init() {

	err := godotenv.Load(".env")

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

	repoPost, repoTag,repoComment := initializeRepository(db)
	servicePost,serviceTag, serviceComment := initializeServices(repoPost, repoTag, repoComment)
	handlerPost,handlerTag, handlerComment := initializeHandlers(servicePost,serviceTag,serviceComment)
	handleFunc(handlerPost,handlerTag,handlerComment)
	fmt.Println(os.Getenv("Port is:"+"PORT"))
}
