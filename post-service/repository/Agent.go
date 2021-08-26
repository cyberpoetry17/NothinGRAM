package repository

import (
	"fmt"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AgentRepo struct {
	Database *gorm.DB
}

func(repo *AgentRepo) CreateAgent(agent *data.Agent)(error,uuid.UUID){
	result := repo.Database.Create(agent)
	if result.Error != nil {
		return result.Error,uuid.Nil
	}
	fmt.Println(result.RowsAffected)
	return nil,agent.ID
}

func (repo *AgentRepo) GetAll() []data.Agent{
	var agents []data.Agent
	repo.Database.Find(&agents)
	return agents
}

func (repo AgentRepo) RemoveAgent(id string) bool{
	var posts = repo.GetAll()
	for _,element := range posts{
		if element.ID.String() == id {
			repo.Database.Delete(&element)
			return true
		}
	}
	return false
}