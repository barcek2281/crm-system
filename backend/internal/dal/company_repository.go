package dal

import (
	"crmsystem/internal/model"
	"database/sql"
)

type CompanyRepo struct {
	db *sql.DB
}

func (j *CompanyRepo) FindCompanyByName(name string) (model.Company, error) {
	q := `SELECT company_id, description FROM company WHERE name = $1`
	company := model.Company{Name: name}
	err := j.db.QueryRow(q, name).Scan(&company.Id, &company.Description)
	if err != nil {
		return model.Company{}, err
	}

	return company, nil
}

func (j *CompanyRepo) CreateJob(company model.Company) (string, error) {
	q := `INSERT INTO job_title (name, description) VALUES ($1, $2) RETURNING job_title_id`
	if err := j.db.QueryRow(q, company.Name, company.Description).Scan(&company.Id); err != nil {
		return "", err
	}
	return company.Id, nil
}
