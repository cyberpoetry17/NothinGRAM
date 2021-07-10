package repository

import (
	"fmt"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"gorm.io/gorm"
)

type ReportedPostRepo struct {
	Database *gorm.DB
}

func (repo ReportedPostRepo) CreateReport(report *data.ReportedPost) error {
	result := repo.Database.Create(report)
	if(result.Error != nil){
		return result.Error
	}
	fmt.Println(result.RowsAffected)
	return nil
}

func (repo ReportedPostRepo) GetAllReportsForPost (postId string) []data.ReportedPost{
	var reports []data.ReportedPost
	var backList []data.ReportedPost
	repo.Database.Find(&reports)
	for _,element := range reports{
		if element.PostId.String() == postId {
			backList = append(backList,element)
		}
	}
	return backList
}

func (repo ReportedPostRepo) CheckIfReportedByUser (report *data.ReportedPost) bool{
	var reports []data.ReportedPost
	reports = repo.GetAllReportsForPost(report.PostId.String())
	for _,element := range reports{
		if element.UserId == report.UserId{
			return true
		}
	}
	return false
}