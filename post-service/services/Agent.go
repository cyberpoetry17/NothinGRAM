package services

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
	"github.com/google/uuid"
)

type AgentService struct {
	AgentRepo *repository.AgentRepo
}

func (service *AgentService) CreateAgent(agent *data.Agent) (error,uuid.UUID){
	return service.AgentRepo.CreateAgent(agent)
}

func (service *AgentService) GetAll() []data.Agent{
	return service.AgentRepo.GetAll()
}

func (service AgentService) RemoveAgent(id string) bool{
	return service.AgentRepo.RemoveAgent(id)
}