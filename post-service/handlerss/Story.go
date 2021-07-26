package handlerss

import (
	"encoding/json"
	"fmt"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/DTO"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
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
	for _,el := range(stories){
		fmt.Println(el)
	}
	json.NewEncoder(w).Encode(stories)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

