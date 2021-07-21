package services

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
)

type FollowerService struct {
	Repo *repository.FollowerRepo
}

func (service *FollowerService) FollowUser(follower *data.Follower) error {
	return service.Repo.FollowUser(follower)
}

func (service *FollowerService) UnfollowUser(follower *data.Follower) error {
	return service.Repo.UnfollowUser(follower)
}

func (service *FollowerService) FollowStatusForProfile(follower *data.Follower) bool {
	return service.Repo.FollowStatusForProfile(follower)
}

func (service *FollowerService) FollowedByUser(userid string) []string {
	return service.Repo.FollowedByUser(userid)
}

// func (service *FollowerService) GetAllFollowersForUser(userId string) ([]data.Follower, error) {
// 	id, err := uuid.Parse(userId)
// 	if err != nil {
// 		print(err)
// 		return nil, err
// 	}
// 	followers := service.Repo.GetAllForOneUser(id)
// 	return followers, nil
// }
