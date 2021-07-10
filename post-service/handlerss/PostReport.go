package handlerss

import (
	"encoding/json"
	"fmt"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/data"
	"github.com/cyberpoetry17/NothinGRAM/UserAPI/services"
	"net/http"
)

type ReportedPostHandler struct {
	Service *services.ReportedPostService
}

func (handler *ReportedPostHandler) CreateReport(w http.ResponseWriter, r *http.Request) {
	fmt.Println("reporting post..")
	var report data.ReportedPost
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.CreateReport(&report)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *ReportedPostHandler) CheckIfUserReportedPost (w http.ResponseWriter,r *http.Request){
	var report data.ReportedPost
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	reported := handler.Service.CheckIfUserReportedPost(&report)

	if reported != false {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(true)
	} else {
		_ = json.NewEncoder(w).Encode(false)

	}
}