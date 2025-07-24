package handler

import (
	"crmsystem/internal/model"
	"crmsystem/internal/service"
	"crmsystem/internal/util"
	"encoding/json"
	"net/http"
)

type Company struct {
	companyService *service.Company
}

func NewCompanyHandler(cs *service.Company) *Company {
	return &Company{
		companyService: cs,
	}
}

func (c *Company) CreateJob() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		company := model.Company{}
		if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
			util.ResponseError(w, http.StatusBadRequest, err)
			return
		}

		id, err := c.companyService.CreateCompany(company)
		if err != nil {
			util.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		company.Id = id
		util.ResponseJSON(w, http.StatusOK, company)
	}
}
