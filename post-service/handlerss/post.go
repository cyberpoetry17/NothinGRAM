package handlerss

import (
	"encoding/json"
	"fmt"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/DTO"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type PostHandler struct {
	Service *services.PostService
}

func (handler *PostHandler) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all posts..")
	json.NewEncoder(w).Encode(handler.Service.GetAllPosts())
	for _,el := range handler.Service.GetAllPosts(){
		fmt.Println("Results: " + el.ID.String())
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func (handler *PostHandler) GetNonPrivatePosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting non private posts..")
	json.NewEncoder(w).Encode(handler.Service.GetNonPrivatePosts())
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func (handler *PostHandler) GetNonPrivatePostsForUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting non private posts for user..")
	vars := mux.Vars(r)
	id := vars["userid"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	posts,err := handler.Service.GetNonPrivatePostsForUser(id)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_ = json.NewEncoder(w).Encode(posts)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func (handler *PostHandler) GetPostsByUserID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all posts for specified user..")
	vars := mux.Vars(r)
	id := vars["userid"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(handler.Service.GetPostsByUserID(id))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func (handler *PostHandler) GetPostsByLocation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all posts by location..")
	vars := mux.Vars(r)
	id := vars["location"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idsplit := strings.Split(id,"@")
	var frontList []data.Post
	for _,element := range idsplit{
		frontList = append(frontList,handler.Service.GetPostsByLocation(strings.TrimSpace(element))...)
	}
	json.NewEncoder(w).Encode(frontList)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostHandler) GetPostsByTags(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all posts by tags..")
	vars := mux.Vars(r)
	id := vars["tag"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idsplit := strings.Split(id," ")
	var frontList []data.Post
	for _,element := range idsplit{
		frontList = append(frontList,handler.Service.GetPostsByTags(strings.TrimSpace(element))...)
	}
	json.NewEncoder(w).Encode(frontList)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func (handler *PostHandler) GetLikedByUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all posts user liked..")
	vars := mux.Vars(r)
	id := vars["userid"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(handler.Service.GetLikedByUser(id))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostHandler) GetDislikedByUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all posts user disliked..")
	vars := mux.Vars(r)
	id := vars["userid"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(handler.Service.GetDislikedByUser(id))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")
	var post DTO.PostDTO
	fmt.Println(post.Description)
	//time2 :=time.Now()
	//fmt.Println(json.NewEncoder(w).Encode(time2))
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(post)
	err = handler.Service.CreatePost(&post)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	fmt.Println("created desc:"+post.Description)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostHandler) Verify(w http.ResponseWriter, r *http.Request) {
	fmt.Println("verifying")
	vars := mux.Vars(r)
	id := vars["picpath"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	exists, err := handler.Service.PostExists(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if exists {
		w.WriteHeader(http.StatusOK)
		fmt.Println("EXISTS")
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (handler *PostHandler) GetTagsForPost (w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	id := vars["postid"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tags,err := handler.Service.GetTagsForPost(id)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_ = json.NewEncoder(w).Encode(tags)
}

func (handler *PostHandler) GetUsernameByPostUserID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting username by id..")
	vars := mux.Vars(r)
	id := vars["userid"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	username := handler.Service.GetUsernameByPostUserID(id)
	if username != "" {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(username)
		fmt.Println("EXISTS")
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (handler *PostHandler) AddTagToPost(w http.ResponseWriter, r *http.Request){
	fmt.Println("creating")
	var postTagDto DTO.PostTagDTO
	err := json.NewDecoder(r.Body).Decode(&postTagDto)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.AddTagToPost(postTagDto.Tag, postTagDto.PostId)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostHandler) AddLocationToPost(w http.ResponseWriter, r *http.Request){
	fmt.Println("creating")
	var postLocationDto DTO.PostLocationDTO
	err := json.NewDecoder(r.Body).Decode(&postLocationDto)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.AddLocationToPost(postLocationDto.Location, postLocationDto.PostId)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	fmt.Println("location put on post")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}