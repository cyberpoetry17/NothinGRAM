package services

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
)

type BlockedService struct {
	Repo         *repository.BlockedRepo
	RepoFollower *repository.FollowerRepo
}

func (service *BlockedService) CreateBlockedUser(blockedUser *data.Blocked) error {
	// followers := service.RepoFollower.GetAll()
	// for i := 0; i < len(followers); i++ {
	// 	if followers[i].IDFollower == blockedUser.BlockedID && followers[i].IDUser == blockedUser.UserID {
	// 		service.RepoFollower.UnfollowUserByIds(blockedUser.BlockedID, blockedUser.UserID) //trebalo bi da radi
	// 	}
	// }

	return service.Repo.CreateBlocked(blockedUser)
}

// func (service *BlockedService) GetAllBlockedUsers(userID string) ([]data.Blocked, error) {

// 	blockedUsers, error := service.Repo.GetAllBlockedUsersByID(userID)
// 	if error != nil {
// 		return nil, error
// 	}
// 	return blockedUsers, nil
// }

func (service *BlockedService) RemoveBlockedUser(blockedUser *data.Blocked) error {
	return service.Repo.RemoveBlocked(blockedUser)
}

func (service *BlockedService) BlockedStatusForProfile(blocked *data.Blocked) bool {
	return service.Repo.BlockStatusForProfile(blocked)
}
