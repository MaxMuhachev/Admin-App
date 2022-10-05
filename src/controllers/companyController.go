package controllers

import (
	"content/src/app"
	"content/src/models"
	"content/src/storage"
	"content/src/utils"
	"encoding/json"
	"net/http"
)

func HandlerApiCompany(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		r.ParseForm()
		var (
			movies        = []models.DataSelect{}
			err           error
			totalElements uint
		)
		search := r.Form.Get("search")
		page := uint(utils.ConvertUint(r.Form.Get("page")))
		movies, err = GetCompanySelect(search, models.DEFAULT_SELECT_LIMIT, page)
		if err == nil {
			totalElements, err = GetSingleResult(storage.GetCompanyCount, search)
			if err == nil {
				res := models.ResponseSelect{Data: movies, TotalElements: totalElements, Page: page, Size: models.DEFAULT_SELECT_LIMIT}
				json.NewEncoder(w).Encode(&res)
				w.WriteHeader(http.StatusOK)
			}
		}
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
		}
	}
}

func GetCompanySelect(search string, limit uint, page uint) ([]models.DataSelect, error) {
	connect := app.NewConnect()

	var res []models.DataSelect

	rows, err := connect.Mysql.Queryx(storage.GetCompanySelect, search, limit, page*limit)
	if err != nil {
		return nil, err
	} else {
		for rows.Next() {
			var company = models.DataSelect{}
			err := rows.StructScan(&company)
			if err != nil {
				return nil, err
			}
			res = append(res, company)
		}
	}

	return res, err
}
