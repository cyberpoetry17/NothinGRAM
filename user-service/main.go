package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"github.com/cyberpoetry17/NothinGRAM/UserAPI/handlers"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
)

func initializeRepository(database *gorm.DB) (*repository.UserRepo, *repository.BlockedRepo, *repository.MutedRepo, *repository.FollowerRepo, *repository.FollowerRequestRepo, *repository.CloseFollowerRepository, *repository.VerificationRequestRepo) {
	return &repository.UserRepo{Database: database}, &repository.BlockedRepo{Database: database}, &repository.MutedRepo{Database: database}, &repository.FollowerRepo{Database: database}, &repository.FollowerRequestRepo{Database: database}, &repository.CloseFollowerRepository{Database: database}, &repository.VerificationRequestRepo{Database: database}
}

func initializeServices(repo *repository.UserRepo, repoBlocked *repository.BlockedRepo, repoMuted *repository.MutedRepo, repoFollower *repository.FollowerRepo, repoFollowerRequest *repository.FollowerRequestRepo, repoCloseFollower *repository.CloseFollowerRepository, repoVerificationRequest *repository.VerificationRequestRepo) (*services.UserService, *services.BlockedService, *services.MutedService, *services.FollowerService, *services.FollowerRequestService, *services.CloseFollowerService, *services.VerificationRequestService) {
	return &services.UserService{Repo: repo, RepoFollower: repoFollower, RepoCloseFollower: repoCloseFollower, MutedRepo: repoMuted, BlockedRepo: repoBlocked}, &services.BlockedService{Repo: repoBlocked}, &services.MutedService{Repo: repoMuted}, &services.FollowerService{Repo: repoFollower, RepoMuted: repoMuted, RepoBlocked: repoBlocked}, &services.FollowerRequestService{Repo: repoFollowerRequest}, &services.CloseFollowerService{Repo: repoCloseFollower}, &services.VerificationRequestService{Repo: repoVerificationRequest}
}

func initializeHandlers(service *services.UserService, serviceBlocked *services.BlockedService, serviceMuted *services.MutedService, serviceFollower *services.FollowerService, serviceFollowerRequest *services.FollowerRequestService, serviceCloseFollower *services.CloseFollowerService, serviceVerificationRequest *services.VerificationRequestService) (*handlers.UserHandler, *handlers.BlockedHandler, *handlers.MutedHandler, *handlers.FollowerHandler, *handlers.FollowerRequestHandler, *handlers.CloseFollowerHandler, *handlers.VerificationRequestHandler) {
	return &handlers.UserHandler{Service: service}, &handlers.BlockedHandler{Service: serviceBlocked}, &handlers.MutedHandler{Service: serviceMuted}, &handlers.FollowerHandler{Service: serviceFollower}, &handlers.FollowerRequestHandler{Service: serviceFollowerRequest}, &handlers.CloseFollowerHandler{Service: serviceCloseFollower, UserServ: service}, &handlers.VerificationRequestHandler{Service: serviceVerificationRequest, ServiceUser: service}
}
func handleFuncUser(handler *handlers.UserHandler, handlerBlocked *handlers.BlockedHandler, handlerMuted *handlers.MutedHandler, followerHandler *handlers.FollowerHandler, handlerFollowerRequest *handlers.FollowerRequestHandler, handlerCloseFollower *handlers.CloseFollowerHandler, handlerVerificationRequest *handlers.VerificationRequestHandler) {
	router := mux.NewRouter().StrictSlash(true)

	// c := cors.New(cors.Options{AllowedOrigins: []string{"*"}, AllowCredentials: true})
	// hand := c.Handler(router)

	router.HandleFunc("/", handler.Hello).Methods("GET")
	router.HandleFunc("/user", handler.GetById).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/register", handler.CreateUser).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/update", handler.UpdateUser).Methods(http.MethodPut, http.MethodOptions)
	router.HandleFunc("/auth", handler.AuthorizationToken).Methods("POST")
	router.HandleFunc("/logout", handler.Logout).Methods("POST")
	router.HandleFunc("/verify/{userId}", handler.Verify).Methods("GET")
	router.HandleFunc("/getpublicuserids", handler.GetPublicUserIds).Methods(http.MethodGet)
	router.HandleFunc("/login", handler.LoginUser).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/username/{usernamebyid}", handler.GetUsernameById).Methods("GET")
	router.HandleFunc("/getuserbyusername/{username}", handler.GetUserByUsernameForProfile).Methods(http.MethodGet)
	router.HandleFunc("/getuseridandprivatebyusername/{username}", handler.GetUserIdByUsernameForProfile).Methods(http.MethodGet)
	router.HandleFunc("/GetUserProfilePrivacy", handler.GetUserProfilePrivacy).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/getuserwhofollow", handler.GetAllUserFollowersById).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/deleteprofile/{userid}", handler.DeleteProfile).Methods(http.MethodPost)
	router.HandleFunc("/confirmagent", handler.CreateUserFromAgent).Methods(http.MethodPost,http.MethodOptions)

	router.HandleFunc("/auth", handler.AuthorizationToken).Methods("POST")

	router.HandleFunc("/block", handlerBlocked.BlockUser).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/unblock", handlerBlocked.UnblockUser).Methods("POST")
	// router.HandleFunc("/allblockedusers/{userID}", handlerBlocked.GetAllBlockedUsers).Methods("GET")
	router.HandleFunc("/getblockedstatus", handlerBlocked.BlockStatusForProfile).Methods(http.MethodPost)

	router.HandleFunc("/createMuted", handlerMuted.CreateMutedUser).Methods(http.MethodPost)
	router.HandleFunc("/removeMuted", handlerMuted.RemoveMutedUser).Methods(http.MethodPost)
	router.HandleFunc("/getmutedstatus", handlerMuted.MutedStatusForProfile).Methods(http.MethodPost)
	// router.HandleFunc("/allmutedusers/{userID}", handlerMuted.GetAllMutedUsers).Methods("GET")

	router.HandleFunc("/follow", followerHandler.FollowUser).Methods(http.MethodPost)
	router.HandleFunc("/getfollowstatus", followerHandler.FollowStatusForProfile).Methods(http.MethodPost)
	router.HandleFunc("/unfollow", followerHandler.UnfollowUser).Methods("POST")
	router.HandleFunc("/getallfollowedforloggeduser/{userid}", followerHandler.FollowedByUser).Methods(http.MethodGet)

	router.HandleFunc("/createfollowrequest", handlerFollowerRequest.CreateFollowRequest).Methods(http.MethodPost)
	router.HandleFunc("/getallrequests/{userid}", handlerFollowerRequest.GetAllRequests).Methods(http.MethodGet)
	router.HandleFunc("/deleterequest/{requestid}", handlerFollowerRequest.DeleteRequest).Methods(http.MethodGet)

	router.HandleFunc("/addclosefollower", handlerCloseFollower.AddCloseFollower).Methods(http.MethodPost)
	router.HandleFunc("/removeclosefollower", handlerCloseFollower.RemoveCloseFollower).Methods(http.MethodPost)
	router.HandleFunc("/setclosefollowers", handlerCloseFollower.ModifyCloseFollowers).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/getclosefollowers", handler.GetAllCloseUserFollowersById).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/getAllCloseFollowersForUser/{userid}", handlerCloseFollower.GetAllCloseFollowerUser).Methods(http.MethodGet)

	router.HandleFunc("/verification", handlerVerificationRequest.CreateVerificationRequest).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/waitlistedverificationrequests", handlerVerificationRequest.GetAllWaitlistedVerificationRequests).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/acceptverification", handlerVerificationRequest.AcceptUserVerificationRequest).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/declineverification", handlerVerificationRequest.DeclineUserVerificationRequest).Methods(http.MethodPost, http.MethodOptions)

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
	loc, err := time.LoadLocation("Europe/Budapest")
	if err != nil {
		println("problem with local zone")
	}
	// handle err
	time.Local = loc //

	repositoryUser, repositoryBlocked, repositoryMuted, repositoryFollower, repoFollowerRequest, repoCloseFollower, repoVerificationRequest := initializeRepository(db)
	serviceUser, serviceBLocked, serviceMuted, serviceFollower, serviceFollowerRequest, serviceCloseFollower, serviceVerificationRequest := initializeServices(repositoryUser, repositoryBlocked, repositoryMuted, repositoryFollower, repoFollowerRequest, repoCloseFollower, repoVerificationRequest)
	handlerUser, handlerBlocked, handlerMuted, handlerFollower, handlerFollowerRequest, handlerCloseFollower, handlerVerificationRequest := initializeHandlers(serviceUser, serviceBLocked, serviceMuted, serviceFollower, serviceFollowerRequest, serviceCloseFollower, serviceVerificationRequest)
	handleFuncUser(handlerUser, handlerBlocked, handlerMuted, handlerFollower, handlerFollowerRequest, handlerCloseFollower, handlerVerificationRequest)
}
