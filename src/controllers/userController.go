package controllers

import (
	"content/src/app"
	"content/src/config"
	"content/src/models"
	"content/src/storage"
	"content/src/utils"
	"encoding/json"
	"net/http"
)

func HandlerApiUsers(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		r.ParseForm()
		var (
			users         []models.DataSelect
			err           error
			totalElements uint
		)
		search := r.Form.Get("search")
		page := uint(utils.ConvertUint(r.Form.Get("page")))
		users, err = GetUsersSelect(search, models.DEFAULT_SELECT_LIMIT, page)
		if err == nil {
			totalElements, err = GetDoubleAttrResult(storage.GetUsersCount, search, search)
			if err == nil {
				res := models.ResponseSelect{Data: users, TotalElements: totalElements, Page: page, Size: models.DEFAULT_SELECT_LIMIT}
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

func GetUsersSelect(search string, limit uint, page uint) ([]models.DataSelect, error) {

	var res = []models.DataSelect{}

	err := app.Conn.Mysql.Select(&res, storage.GetUsersSelect, search, search, limit, page*limit)
	if err != nil {
		return nil, err
	}
	return res, err
}

func GetDoubleAttrResult(query string, search string, search2 string) (uint, error) {
	var res uint

	err := app.Conn.Mysql.Get(&res, query, search, search2)
	if err != nil {
		return 0, err
	}
	return res, err
}

func GetUserByEmailPassword(email string, password string) (*models.User, error) {
	var user = &models.User{}

	err := app.Conn.Mysql.Get(user, storage.GetUserByEmailPassword, email, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func HandlerEditUser(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieUserWriteHeader(w, r) {
		userEmail, _ := r.Cookie(config.USER_PERSMISSION)

		user, err := GetUserById(userEmail.Value)
		app.RenderTemplate(w, "user/edit/content-edit-user", &app.Page{Title: utils.USER, User: user}, &err)
	}
}

func GetUserById(email string) (*models.User, error) {
	var res = &models.User{}

	err := app.Conn.Mysql.Get(res, storage.GetUserByEmail, email)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func HandlerEditPostUser(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieUserWriteHeader(w, r) {
		user := ParseUserForm(r)
		oldEmail := r.Form.Get("oldEmail")

		_, err := app.Conn.Mysql.Query(
			storage.UpdateUserByEmail,
			user.Email,
			user.FIO,
			user.Login,
			user.Password,
			user.Floor,
			oldEmail,
		)

		if !utils.ThrowError(err, w) {
			utils.ClearTokenHandler(w, config.USER_PERSMISSION)
			utils.SetCookie(w, config.USER_PERSMISSION, user.Email)
			app.RenderTemplate(w, "user/edit/content-edit-user", &app.Page{Title: utils.USER, User: user, Success: utils.USER_SAVED}, &err)
		}

	}
}

func ParseUserForm(r *http.Request) *models.User {
	r.ParseForm()
	fio := r.Form.Get("FIO")
	email := r.Form.Get("email")
	floor := r.Form.Get("floor")
	login := r.Form.Get("login")
	password := r.Form.Get("password")

	var employee = models.User{}
	employee.FIO = fio
	employee.Email = email
	employee.Login = login
	employee.Password = password
	employee.Floor = floor
	return &employee
}
