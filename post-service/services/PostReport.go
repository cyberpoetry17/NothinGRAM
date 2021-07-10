package services

import (
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/repository"
)

type ReportedPostService struct {
	Repo *repository.ReportedPostRepo
}

func (service *ReportedPostService) CreateReport (report *data.ReportedPost) error{
	return service.Repo.CreateReport(report)
}

func (service *ReportedPostService) CheckIfUserReportedPost (report *data.ReportedPost) bool{
	return service.Repo.CheckIfReportedByUser(report)
}