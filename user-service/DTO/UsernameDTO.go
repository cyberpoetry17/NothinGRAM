package DTO

type UsernameDTO struct {
	Usernames        []string ` json:"usernames" `
	RemovedUsernames []string `json:"removedusernames"`
}
