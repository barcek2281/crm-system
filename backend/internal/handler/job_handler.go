package handler

import (
	"crmsystem/internal/model"
	"crmsystem/internal/service"
	"crmsystem/internal/util"
	"encoding/json"
	"net/http"
)

type Job struct {
	jobService *service.Job
}

func NewJobHandler(job *service.Job) *Job {
	return &Job{
		jobService: job,
	}
}

func (j *Job) CreateJob() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		job := model.Job{}
		if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
			util.ResponseError(w, http.StatusBadRequest, err)
			return
		}

		id, err := j.jobService.CreateJob(job)
		if err != nil {
			util.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		job.JobId = id
		util.ResponseJSON(w, http.StatusOK, job)
	}
}