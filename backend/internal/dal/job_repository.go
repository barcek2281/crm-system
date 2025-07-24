package dal

import (
	"crmsystem/internal/model"
	"database/sql"
)

type JobRepo struct {
	db *sql.DB
}

func (j *JobRepo) FindJobByName(name string) (model.Job, error) {
	q := `SELECT job_title_id FROM job_title WHERE name = $1`
	id := ""
	err := j.db.QueryRow(q, name).Scan(&id)
	if err != nil {
		return model.Job{}, err
	}

	job := model.Job{
		Name:  name,
		JobId: id,
	}
	return job, nil
}


func (j *JobRepo) CreateJob(name string) (string, error) {
	q := `INSERT INTO job_title (name) VALUES ($1) RETURNING job_title_id`
	id := ""
	if err := j.db.QueryRow(q, name).Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}