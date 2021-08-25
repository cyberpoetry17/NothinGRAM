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

func initializeRepository(database *gorm.DB) (*repository.PostRepo,*repository.TagRepo,*repository.CommentRepo,*repository.LikeRepo,*repository.DislikeRepo, *repository.MediaRepo,*repository.LocationRepo,*repository.ReportedPostRepo,*repository.StoryRepo) {
	return &repository.PostRepo{Database: database}, &repository.TagRepo{Database: database}, &repository.CommentRepo{Database: database},&repository.LikeRepo{Database: database},&repository.DislikeRepo{Database: database}, &repository.MediaRepo{Database: database},&repository.LocationRepo{Database: database},&repository.ReportedPostRepo{Database: database}, &repository.StoryRepo{Database: database}
}

func initializeServices(repoPost *repository.PostRepo, repoTag *repository.TagRepo, repoComment *repository.CommentRepo, repoLike *repository.LikeRepo,repoDislike *repository.DislikeRepo, repoMedia *repository.MediaRepo,repoLocation *repository.LocationRepo,repoReport *repository.ReportedPostRepo,repoStory *repository.StoryRepo) (*services.PostService,*services.TagService,*services.CommentService,*services.LikeService,*services.DislikeService,*services.MediaService,*services.LocationService,*services.ReportedPostService,*services.StoryService) {
	return &services.PostService{PostRepo: repoPost,TagRepo: repoTag,LikeRepo:repoLike,DislikeRepo: repoDislike,MediaRepo: repoMedia,LocationRepo: repoLocation,ReportRepo: repoReport,CommentRepo: repoComment}, &services.TagService{Repo: repoTag}, &services.CommentService{Repo: repoComment},&services.LikeService{Repo: repoLike},&services.DislikeService{Repo: repoDislike},&services.MediaService{Repo: repoMedia},&services.LocationService{Repo: repoLocation},&services.ReportedPostService{Repo: repoReport}, &services.StoryService{StoryRepo: repoStory,PostRepo: repoPost, MediaRepo: repoMedia}
}

func initializeHandlers(servicePost *services.PostService,serviceTag *services.TagService, serviceComment *services.CommentService,serviceLike *services.LikeService,serviceDislike *services.DislikeService, serviceMedia *services.MediaService,serviceLocation *services.LocationService,serviceReport *services.ReportedPostService, serviceStory *services.StoryService) (*handlerss.PostHandler,*handlerss.TagHandler,*handlerss.CommentHandler,*handlerss.LikeHandler,*handlerss.DislikeHandler,*handlerss.MediaHandler,*handlerss.LocationHandler,*handlerss.ReportedPostHandler,*handlerss.StoryHandler) {
	return &handlerss.PostHandler{Service: servicePost}, &handlerss.TagHandler{Service: serviceTag}, &handlerss.CommentHandler{Service: serviceComment},&handlerss.LikeHandler{Service: serviceLike},&handlerss.DislikeHandler{Service: serviceDislike}, &handlerss.MediaHandler{Service: serviceMedia},&handlerss.LocationHandler{Service: serviceLocation},&handlerss.ReportedPostHandler{Service: serviceReport},&handlerss.StoryHandler{Service: serviceStory}
}
func handleFunc(handler *handlerss.PostHandler,tagHandler *handlerss.TagHandler, commentHandler *handlerss.CommentHandler,likeHandler *handlerss.LikeHandler,dislikeHandler *handlerss.DislikeHandler, mediaHandler *handlerss.MediaHandler,locationHandler *handlerss.LocationHandler,reportHandler *handlerss.ReportedPostHandler, storyHandler *handlerss.StoryHandler) {
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
	reportHandleFuncs(reportHandler,router)
	storyHandleFuncs(router, storyHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("POST_SERVICE_PORT")), hand))
}

func storyHandleFuncs(router *mux.Router, storyHandler *handlerss.StoryHandler) {
	router.HandleFunc("/addStory", storyHandler.CreateStory).Methods("POST")
	router.HandleFunc("/getAllStories", storyHandler.GetAllActiveStories).Methods("GET")
	router.HandleFunc("/getUserStories/{userId}",storyHandler.GetAllUserStories).Methods("GET")
	router.HandleFunc("/GetCloseFrinedStoriesForUser/{userId}",storyHandler.GetCloseFrinedStoriesForUser).Methods("GET")
	router.HandleFunc("/AddToStoryHighlights/{storyId}",storyHandler.AddToStoryHighlights).Methods("POST")
	router.HandleFunc("/RemoveFromStoryHighlights/{storyId}",storyHandler.RemoveFromStoryHighlights).Methods("POST")
	router.HandleFunc("/GetAllStoryHighlights/{userId}",storyHandler.GetAllStoryHighlights).Methods("GET")
	router.HandleFunc("/GetActiveStoriesByUserId/{userId}",storyHandler.GetActiveStoriesByUserId).Methods("GET")
}

func mediaHandleFuncs(router *mux.Router, mediaHandler *handlerss.MediaHandler) {
	router.HandleFunc("/addMedia/", mediaHandler.CreateMedia).Methods("POST")
	router.HandleFunc("/EditMedia/", mediaHandler.EditMedia).Methods("POST")
	router.HandleFunc("RemoveMedia/", mediaHandler.RemoveMedia).Methods("DELETE")
	router.HandleFunc("/GetMediaForPost",mediaHandler.GetMediaForPost).Methods("GET")
	router.HandleFunc("/GetMediaForStory",mediaHandler.GetMediaForStory).Methods("GET")
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
	router.HandleFunc("/deletedislike",dislikeHandler.DeleteDislike).Methods("POST")
}
func likeHandleFuncs(router *mux.Router, likeHandler *handlerss.LikeHandler) {
	router.HandleFunc("/alllikesforpost/{postid}", likeHandler.GetAllLikesForPost).Methods("GET")
	router.HandleFunc("/createlike", likeHandler.CreateLike).Methods("POST")
	router.HandleFunc("/checkiflikedbyuser",likeHandler.CheckIfUserLikedPost).Methods("POST")
	router.HandleFunc("/deletelike",likeHandler.DeleteLike).Methods("POST")
}

func commentHandleFuncs(router *mux.Router, commentHandler *handlerss.CommentHandler) {
	router.HandleFunc("/addComment/", commentHandler.CreateComment).Methods("POST")
	router.HandleFunc("/editComment/", commentHandler.EditComment).Methods("POST")
	router.HandleFunc("/removeComment/", commentHandler.DeleteComment).Methods("DELETE")
	router.HandleFunc("/getcommentsforpost/{postid}",commentHandler.GetAllByPostId).Methods("GET")
}

func tagHandleFuncs(router *mux.Router, tagHandler *handlerss.TagHandler) {
	router.HandleFunc("/addTag/", tagHandler.CreateTag).Methods("POST")
	router.HandleFunc("/editTag/", tagHandler.EditTag).Methods("POST")
	router.HandleFunc("/removeTag/", tagHandler.DeleteTag).Methods("DELETE")
	router.HandleFunc("/filterpublicmaterialbytagid/{tagid}",tagHandler.FilterPublicMaterialByTagId).Methods("GET")
	router.HandleFunc("/getAllTagsNames",tagHandler.GetAllTagNames).Methods("GET")
	router.HandleFunc("/getAllTags",tagHandler.GetAllTags).Methods("GET")
}

func postHandleFuncs(handler *handlerss.PostHandler, router *mux.Router) {
	router.HandleFunc("/getnonprivateposts", handler.GetNonPrivatePosts).Methods("GET")
	router.HandleFunc("/getnonprivateposts/{userid}", handler.GetNonPrivatePostsForUser).Methods("GET")
	router.HandleFunc("/", handler.Hello).Methods("GET")
	router.HandleFunc("/createpost", handler.CreatePost).Methods("POST")
	router.HandleFunc("/allpostsbyuserid/{userid}",handler.GetPostsByUserID).Methods("GET")
	router.HandleFunc("/verify/{description}", handler.Verify).Methods("GET")
	router.HandleFunc("/addTagToPost", handler.AddTagToPost).Methods("POST")
	router.HandleFunc("/addlocationtopost", handler.AddLocationToPost).Methods("POST")
	router.HandleFunc("/getusernamebyid/{userid}", handler.GetUsernameByPostUserID).Methods("GET")
	router.HandleFunc("/getlikedbyuser/{userid}", handler.GetLikedByUser).Methods("GET")
	router.HandleFunc("/getdislikedbyuser/{userid}", handler.GetDislikedByUser).Methods("GET")
	router.HandleFunc("/tagsforpost/{postid}", handler.GetTagsForPost).Methods("GET")
	router.HandleFunc("/postsbylocation/{location}", handler.GetPostsByLocation).Methods("GET")
	router.HandleFunc("/postsbytags/{tag}", handler.GetPostsByTags).Methods("GET")
	router.HandleFunc("/getallreported", handler.GetAllReported).Methods("GET")
	router.HandleFunc("/deletepost/{postid}", handler.DeletePost).Methods("POST")
}

func reportHandleFuncs(handler *handlerss.ReportedPostHandler, router *mux.Router) {
	router.HandleFunc("/reportpost", handler.CreateReport).Methods("POST")
	router.HandleFunc("/checkifreportedbyuser",handler.CheckIfUserReportedPost).Methods("POST")
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

	repoPost, repoTag,repoComment,repoLike,repoDislike,repoMedia,repoLocation,repoReport,repoStory := initializeRepository(db)
	servicePost,serviceTag, serviceComment, serviceLike, serviceDislike,serviceMedia,serviceLocation,serviceReport,serviceStory := initializeServices(repoPost, repoTag, repoComment,repoLike,repoDislike,repoMedia,repoLocation,repoReport,repoStory)
	handlerPost,handlerTag, handlerComment,handlerLike, handlerDislike, handlerMedia,handlerLocation,handlerReport,handlerStory := initializeHandlers(servicePost,serviceTag,serviceComment,serviceLike,serviceDislike,serviceMedia,serviceLocation,serviceReport,serviceStory)
	handleFunc(handlerPost,handlerTag,handlerComment,handlerLike,handlerDislike,handlerMedia,handlerLocation,handlerReport,handlerStory)
	fmt.Println(os.Getenv("Port is:"+"PORT"))


}
