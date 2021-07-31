package handlerss

import (
	"encoding/json"
	"fmt"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/DTO"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type StoryHandler struct {
	Service *services.StoryService
}

func (handler *StoryHandler) CreateStory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating story")
	var storyDto DTO.StoryMediaDTO
	err := json.NewDecoder(r.Body).Decode(&storyDto)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.CreateStory(&storyDto)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *StoryHandler) GetAllActiveStories(w http.ResponseWriter, r *http.Request) {
	stories := handler.Service.GetAllActiveStories()
	json.NewEncoder(w).Encode(stories)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *StoryHandler) GetAllUserStories(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["userId"]
	userId,err := uuid.Parse(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	stories,err := handler.Service.GetAllUserStories(userId)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(stories)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *StoryHandler) GetCloseFrinedStoriesForUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["userId"]
	userId,err := uuid.Parse(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	stories := handler.Service.GetCloseFrinedStoriesForUser(userId)
	json.NewEncoder(w).Encode(stories)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *StoryHandler) AddToStoryHighlights(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["storyId"]
	storyId,err := uuid.Parse(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.AddToStoryHighlights(storyId)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *StoryHandler) RemoveFromStoryHighlights(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["storyId"]
	storyId,err := uuid.Parse(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.RemoveFromStoryHighlights(storyId)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *StoryHandler) GetAllStoryHighlights(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["userId"]
	userId,err := uuid.Parse(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	stories := handler.Service.GetAllStoryHighlights(userId)
	json.NewEncoder(w).Encode(stories)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}