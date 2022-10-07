package controllers

import (
	"content/src/app"
	"content/src/config"
	"content/src/models"
	"content/src/utils"
	"errors"
	"net/http"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdmin(r) {
		utils.ClearTokenHandler(w, config.MANAGER_PERSMISSION)
	}
	app.RenderTemplate(w, "login", &app.Page{Title: utils.HELLO + "Челоовек!"}, nil)
}

func HandlerPostIndex(w http.ResponseWriter, r *http.Request) {
	password, email := getFormParams(r)

	employee := GetEmployeeByEmailPassword(email, password)
	if employee != nil {
		utils.SetCookie(w, config.MANAGER_PERSMISSION, email)
		app.RenderTemplate(w, "index", &app.Page{Title: utils.HELLO + employee.FIO + "!"}, nil)
	} else {
		getUserOrError(w, email, password)
	}
}

func getUserOrError(w http.ResponseWriter, email string, password string) {
	var user *models.User
	user, err := GetUserByEmailPassword(email, password)
	if user != nil {
		utils.SetCookie(w, config.USER_PERSMISSION, email)
		app.RenderTemplate(w, "index-user", &app.Page{Title: utils.HELLO + user.FIO + "!"}, &err)
	} else {
		if err == nil || err.Error() == utils.NO_ROWS_RESULT_SET {
			err = errors.New(utils.LOGIN_OR_PASSWORD_NOT_RIGHT)
		}
		app.RenderTemplate(w, "error", &app.Page{Error: &err}, nil)
	}
}

func getFormParams(r *http.Request) (string, string) {
	r.ParseForm()
	password := r.Form.Get("password")
	email := r.Form.Get("email")
	return password, email
}
