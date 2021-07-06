package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"gorm.io/gorm"

	"github.com/cyberpoetry17/NothinGRAM/UserAPI/handlerss"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
)

func initializeRepository(database *gorm.DB) (*repository.PostRepo,*repository.TagRepo,*repository.CommentRepo,*repository.LikeRepo,*repository.DislikeRepo, *repository.MediaRepo,*repository.LocationRepo) {
	return &repository.PostRepo{Database: database}, &repository.TagRepo{Database: database}, &repository.CommentRepo{Database: database},&repository.LikeRepo{Database: database},&repository.DislikeRepo{Database: database}, &repository.MediaRepo{Database: database},&repository.LocationRepo{Database: database}
}

func initializeServices(repoPost *repository.PostRepo, repoTag *repository.TagRepo, repoComment *repository.CommentRepo, repoLike *repository.LikeRepo,repoDislike *repository.DislikeRepo, repoMedia *repository.MediaRepo,repoLocation *repository.LocationRepo) (*services.PostService,*services.TagService,*services.CommentService,*services.LikeService,*services.DislikeService,*services.MediaService,*services.LocationService) {
	return &services.PostService{PostRepo: repoPost,TagRepo: repoTag}, &services.TagService{Repo: repoTag}, &services.CommentService{Repo: repoComment},&services.LikeService{Repo: repoLike},&services.DislikeService{Repo: repoDislike},&services.MediaService{Repo: repoMedia},&services.LocationService{Repo: repoLocation}
}

func initializeHandlers(servicePost *services.PostService,serviceTag *services.TagService, serviceComment *services.CommentService,serviceLike *services.LikeService,serviceDislike *services.DislikeService, serviceMedia *services.MediaService,serviceLocation *services.LocationService) (*handlerss.PostHandler,*handlerss.TagHandler,*handlerss.CommentHandler,*handlerss.LikeHandler,*handlerss.DislikeHandler,*handlerss.MediaHandler,*handlerss.LocationHandler) {
	return &handlerss.PostHandler{Service: servicePost}, &handlerss.TagHandler{Service: serviceTag}, &handlerss.CommentHandler{Service: serviceComment},&handlerss.LikeHandler{Service: serviceLike},&handlerss.DislikeHandler{Service: serviceDislike}, &handlerss.MediaHandler{Service: serviceMedia},&handlerss.LocationHandler{Service: serviceLocation}
}
func handleFunc(handler *handlerss.PostHandler,tagHandler *handlerss.TagHandler, commentHandler *handlerss.CommentHandler,likeHandler *handlerss.LikeHandler,dislikeHandler *handlerss.DislikeHandler, mediaHandler *handlerss.MediaHandler,locationHandler *handlerss.LocationHandler) {
	router := mux.NewRouter().StrictSlash(true)

	c:=cors.New(cors.Options{AllowedOrigins: []string{"*"},AllowCredentials: true})
	hand := c.Handler(router)

	postHandleFuncs(handler, router)
	tagHandleFuncs(router, tagHandler)
	commentHandleFuncs(router, commentHandler)
	likeHandleFuncs(router, likeHandler)
	mediaHandleFuncs(router, mediaHandler)
	dislikeHandleFuncs(router, dislikeHandler)
	locationHandleFuncs(router, locationHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("POST_SERVICE_PORT")), hand))
}

func mediaHandleFuncs(router *mux.Router, mediaHandler *handlerss.MediaHandler) {
	router.HandleFunc("/addMedia/", mediaHandler.CreateMedia).Methods("POST")
	router.HandleFunc("/EditMedia/", mediaHandler.EditMedia).Methods("POST")
	router.HandleFunc("RemoveMedia/", mediaHandler.RemoveMedia).Methods("DELETE")
}

func locationHandleFuncs(router *mux.Router, locationHandler *handlerss.LocationHandler) {
	router.HandleFunc("/locationforpost/{postid}", locationHandler.GetLocationForPost).Methods("GET")
	router.HandleFunc("/createlocation", locationHandler.CreateLocation).Methods("POST")
	router.HandleFunc("/filterpublicmaterialbylocationid/{locationid}",locationHandler.FilterPublicMaterialByLocationId).Methods("GET")
}

func dislikeHandleFuncs(router *mux.Router, dislikeHandler *handlerss.DislikeHandler) {
	router.HandleFunc("/alldislikesforpost/{postid}", dislikeHandler.GetAllDislikesForPost).Methods("GET")
	router.HandleFunc("/createdislike", dislikeHandler.CreateDislike).Methods("POST")
	router.HandleFunc("/checkifdislikedbyuser",dislikeHandler.CheckIfUserDislikedPost).Methods("POST")
}
func likeHandleFuncs(router *mux.Router, likeHandler *handlerss.LikeHandler) {
	router.HandleFunc("/alllikesforpost/{postid}", likeHandler.GetAllLikesForPost).Methods("GET")
	router.HandleFunc("/createlike", likeHandler.CreateLike).Methods("POST")
	router.HandleFunc("/checkiflikedbyuser",likeHandler.CheckIfUserLikedPost).Methods("POST")
}

func commentHandleFuncs(router *mux.Router, commentHandler *handlerss.CommentHandler) {
	router.HandleFunc("/addComment/", commentHandler.CreateComment).Methods("POST")
	router.HandleFunc("/editComment/", commentHandler.EditComment).Methods("POST")
	router.HandleFunc("/removeComment/", commentHandler.DeleteComment).Methods("DELETE")
}

func tagHandleFuncs(router *mux.Router, tagHandler *handlerss.TagHandler) {
	router.HandleFunc("/addTag/", tagHandler.CreateTag).Methods("POST")
	router.HandleFunc("/editTag/", tagHandler.EditTag).Methods("POST")
	router.HandleFunc("/removeTag/", tagHandler.DeleteTag).Methods("DELETE")
	router.HandleFunc("/filterpublicmaterialbytagid/{tagid}",tagHandler.FilterPublicMaterialByTagId).Methods("GET")
}

func postHandleFuncs(handler *handlerss.PostHandler, router *mux.Router) {
	router.HandleFunc("/", handler.Hello).Methods("GET")
	router.HandleFunc("/createpost", handler.CreatePost).Methods("POST")
	router.HandleFunc("/verify/{description}", handler.Verify).Methods("GET")
	router.HandleFunc("/addTagToPost", handler.AddTagToPost).Methods("POST")
	router.HandleFunc("/addlocationtopost", handler.AddLocationToPost).Methods("POST")
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

	repoPost, repoTag,repoComment,repoLike,repoDislike,repoMedia,repoLocation := initializeRepository(db)
	servicePost,serviceTag, serviceComment, serviceLike, serviceDislike,serviceMedia,serviceLocation := initializeServices(repoPost, repoTag, repoComment,repoLike,repoDislike,repoMedia,repoLocation)
	handlerPost,handlerTag, handlerComment,handlerLike, handlerDislike, handlerMedia,handlerLocation := initializeHandlers(servicePost,serviceTag,serviceComment,serviceLike,serviceDislike,serviceMedia,serviceLocation)
	handleFunc(handlerPost,handlerTag,handlerComment,handlerLike,handlerDislike,handlerMedia,handlerLocation)
	fmt.Println(os.Getenv("Port is:"+"PORT"))


}
