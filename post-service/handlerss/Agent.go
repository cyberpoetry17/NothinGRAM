package handlerss

import (
	"encoding/json"
	"fmt"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
	"github.com/gorilla/mux"
	"net/http"
)

type AgentHandler struct {
	Service *services.AgentService
}

func (handler *AgentHandler) CreateAgent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating agent")
	var agent data.Agent
	err := json.NewDecoder(r.Body).Decode(&agent)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(agent)
	err,_ = handler.Service.CreateAgent(&agent)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *AgentHandler) DeleteAgent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting agent")
	vars := mux.Vars(r)
	id := vars["agentemail"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	isItDeleted:= handler.Service.RemoveAgent(id)
	if isItDeleted == false{
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *AgentHandler) GetAll (w http.ResponseWriter,r *http.Request){
	agents := handler.Service.GetAll()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(agents)
}
