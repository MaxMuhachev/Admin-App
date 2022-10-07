package controllers

import (
	"content/src/app"
	"content/src/models"
	"content/src/storage"
	"content/src/utils"
	"encoding/json"
	"net/http"
)

func HandlerStatistic(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		app.RenderTemplate(w, "statistic/content-statistic", &app.Page{}, nil)
	}
}

func HandlerStatisticFilm(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		app.RenderTemplate(w, "statistic/films/content-statistic-films", &app.Page{}, nil)
	}
}

func HandlerStatisticFilmWithDate(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		r.ParseForm()
		startDate := r.Form.Get("startDate")
		endDate := r.Form.Get("endDate")

		var res uint

		err := app.Conn.Mysql.Get(&res, storage.GetStatisticMovie, startDate, endDate)
		if err != nil {
			utils.ThrowError(err, w)
		} else {
			w.Write([]byte(utils.ConvertToString(uint8(res))))
			w.WriteHeader(http.StatusOK)
		}
	}
}

func HandlerStatisticUsers(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		app.RenderTemplate(w, "statistic/users/content-statistic-users", &app.Page{}, nil)
	}
}

func HandlerStatisticUsersWithDate(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		r.ParseForm()
		startDate := r.Form.Get("startDate")
		endDate := r.Form.Get("endDate")

		var res uint

		err := app.Conn.Mysql.Get(&res, storage.GetStatisticUsers, startDate, endDate)
		if err != nil {
			utils.ThrowError(err, w)
		} else {
			w.Write([]byte(utils.ConvertToString(uint8(res))))
			w.WriteHeader(http.StatusOK)
		}
	}
}

func HandlerStatisticComments(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		app.RenderTemplate(w, "statistic/comments/content-statistic-comments", &app.Page{}, nil)
	}
}

func HandlerStatisticCommentsWithDate(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		r.ParseForm()
		startDate := r.Form.Get("startDate")
		endDate := r.Form.Get("endDate")

		var res []*models.CommentStatistic

		err := app.Conn.Mysql.Select(&res, storage.GetStatisticComments, startDate, endDate)
		if err != nil {
			utils.ThrowError(err, w)
		} else {
			json.NewEncoder(w).Encode(res)
			w.WriteHeader(http.StatusOK)
		}
	}
}
