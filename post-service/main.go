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

func initializeRepository(database *gorm.DB) (*repository.PostRepo,*repository.TagRepo,*repository.CommentRepo,*repository.LikeRepo,*repository.DislikeRepo) {
	return &repository.PostRepo{Database: database}, &repository.TagRepo{Database: database}, &repository.CommentRepo{Database: database},&repository.LikeRepo{Database: database},&repository.DislikeRepo{Database: database}
}

func initializeServices(repoPost *repository.PostRepo, repoTag *repository.TagRepo, repoComment *repository.CommentRepo, repoLike *repository.LikeRepo,repoDislike *repository.DislikeRepo) (*services.PostService,*services.TagService,*services.CommentService,*services.LikeService,*services.DislikeService) {
	return &services.PostService{Repo: repoPost}, &services.TagService{Repo: repoTag}, &services.CommentService{Repo: repoComment},&services.LikeService{Repo: repoLike},&services.DislikeService{Repo: repoDislike}
}

func initializeHandlers(servicePost *services.PostService,serviceTag *services.TagService, serviceComment *services.CommentService,serviceLike *services.LikeService,serviceDislike *services.DislikeService) (*handlers.PostHandler,*handlers.TagHandler,*handlers.CommentHandler,*handlers.LikeHandler,*handlers.DislikeHandler) {
	return &handlers.PostHandler{Service: servicePost}, &handlers.TagHandler{Service: serviceTag}, &handlers.CommentHandler{Service: serviceComment},&handlers.LikeHandler{Service: serviceLike},&handlers.DislikeHandler{Service: serviceDislike}
}
func handleFunc(handler *handlers.PostHandler,tagHandler *handlers.TagHandler, commentHandler *handlers.CommentHandler,likeHandler *handlers.LikeHandler,dislikeHandler *handlers.DislikeHandler) {
	router := mux.NewRouter().StrictSlash(true)

	postHandleFuncs(handler, router)
	tagHandleFuncs(router, tagHandler)
	commentHandleFuncs(router, commentHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("POST_SERVICE_PORT")), router))
}

func commentHandleFuncs(router *mux.Router, commentHandler *handlers.CommentHandler) {
	router.HandleFunc("/addComment/", commentHandler.CreateComment).Methods("POST")
	router.HandleFunc("/editComment/", commentHandler.EditComment).Methods("POST")
	router.HandleFunc("/removeComment/", commentHandler.DeleteComment).Methods("DELETE")
}

func tagHandleFuncs(router *mux.Router, tagHandler *handlers.TagHandler) {
	router.HandleFunc("/addTag/", tagHandler.CreateTag).Methods("POST")
	router.HandleFunc("/editTag/", tagHandler.EditTag).Methods("POST")
	router.HandleFunc("/removeTag/", tagHandler.DeleteTag).Methods("DELETE")
}

func postHandleFuncs(handler *handlers.PostHandler, router *mux.Router) {
	router.HandleFunc("/", handler.Hello).Methods("GET")
	router.HandleFunc("/createpost", handler.CreatePost).Methods("POST")
	router.HandleFunc("/verify/{description}", handler.Verify).Methods("GET")
	router.HandleFunc("/addTagToPost", handler.AddTagToPost).Methods("POST")
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

	repoPost, repoTag,repoComment,repoLike,repoDislike := initializeRepository(db)
	servicePost,serviceTag, serviceComment, serviceLike, serviceDislike := initializeServices(repoPost, repoTag, repoComment,repoLike,repoDislike)
	handlerPost,handlerTag, handlerComment,handlerLike, handlerDislike := initializeHandlers(servicePost,serviceTag,serviceComment,serviceLike,serviceDislike)
	handleFunc(handlerPost,handlerTag,handlerComment,handlerLike,handlerDislike)
	fmt.Println(os.Getenv("Port is:"+"PORT"))
}
