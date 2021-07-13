package handlerss

import (
	"encoding/json"
	"fmt"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
	"github.com/gorilla/mux"
	"net/http"
)

type CommentHandler struct {
	Service *services.CommentService
}

func (handler *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating comment")
	var comment data.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(comment)
	err = handler.Service.CreateComment(&comment)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *CommentHandler) EditComment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("edit comment")
	var comment data.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(comment)
	err = handler.Service.EditComment(&comment)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *CommentHandler) GetAllByPostId (w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	id := vars["postid"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	comments,err := handler.Service.GetAllByPostId(id)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(comments)!=0 {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(comments)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (handler *CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete comment")
	var comment data.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(comment)
	err = handler.Service.RemoveComment(&comment)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}