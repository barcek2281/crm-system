package service

import (
	"crmsystem/internal/dal"
	"crmsystem/internal/model"
)

type Company struct {
	companyRepo *dal.CompanyRepo
}

func NewCompanyService(company *dal.CompanyRepo) *Company {
	return &Company{
		companyRepo: company,
	}
}

func (c *Company) CreateCompany(company model.Company) (string, error) {
	return c.companyRepo.CreateCompany(company)
}


