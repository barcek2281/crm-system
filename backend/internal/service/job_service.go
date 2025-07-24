package service

import (
	"crmsystem/internal/dal"
	"crmsystem/internal/model"
)

type Job struct {
	jobRepo *dal.JobRepo
}

func NewJobService(jobRepo *dal.JobRepo) *Job {
	return &Job{
		jobRepo: jobRepo,
	}
}

func (j *Job) CreateJob(job model.Job) (string, error) {
	return j.jobRepo.CreateJob(job.Name)
}
