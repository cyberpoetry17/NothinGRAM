package services

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
)

type FollowerService struct {
	Repo        *repository.FollowerRepo
	RepoBlocked *repository.BlockedRepo
	RepoMuted   *repository.MutedRepo
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
	muted := service.RepoMuted.GetAllMutedUsersByID(userid)
	follower := service.Repo.FollowedByUser(userid)
	var modifiedFollower = follower
	for i := 0; i < len(muted); i++ {
		for j := 0; j < len(follower); j++ {
			if muted[i] == follower[j] {
				modifiedFollower = append(modifiedFollower[:j], modifiedFollower[j+1:]...)
			}

		}

	}

	return modifiedFollower
	// return service.Repo.FollowedByUser(userid)
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
