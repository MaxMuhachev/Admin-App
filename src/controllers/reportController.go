package controllers

import (
	"content/src/app"
	"content/src/models"
	"content/src/storage"
	"content/src/utils"
	"encoding/json"
	"net/http"
)

func HandlerReport(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		app.RenderTemplate(w, "reports/content-report", &app.Page{}, nil)
	}
}

func HandlerReportFilm(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		app.RenderTemplate(w, "reports/films/content-report-films", &app.Page{}, nil)
	}
}

func HandlerReportFilmWithDate(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		r.ParseForm()
		startDate := r.Form.Get("startDate")
		endDate := r.Form.Get("endDate")

		connect := app.NewConnect()

		var res []*models.MovieReport

		rows, err := connect.Mysql.Queryx(storage.GetReportMovie, startDate, endDate)
		if err != nil {
			utils.ThrowError(err, w)
		} else {
			for rows.Next() {
				var movieStatistic = models.MovieReport{}
				err := rows.StructScan(&movieStatistic)
				if err != nil {
					utils.ThrowError(err, w)
				}
				res = append(res, &movieStatistic)
			}
		}
		app.CloseConnect(connect)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}

func HandlerReportUsers(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		app.RenderTemplate(w, "reports/users/content-report-users", &app.Page{}, nil)
	}
}

func HandlerReportUsersWithDate(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		r.ParseForm()
		startDate := r.Form.Get("startDate")
		endDate := r.Form.Get("endDate")

		connect := app.NewConnect()

		var res []*models.UsersReport

		rows, err := connect.Mysql.Queryx(storage.GetReportUsers, startDate, endDate)
		if err != nil {
			utils.ThrowError(err, w)
		} else {
			for rows.Next() {
				var usersReport = models.UsersReport{}
				err := rows.StructScan(&usersReport)
				if err != nil {
					utils.ThrowError(err, w)
				}
				res = append(res, &usersReport)
			}
		}
		app.CloseConnect(connect)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(res)
		}
	}
}

func HandlerReportComments(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		app.RenderTemplate(w, "reports/comments/content-report-comments", &app.Page{}, nil)
	}
}

func HandlerReportCommentsByMovieUser(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		r.ParseForm()
		movieId := r.Form.Get("movie")
		userFio := r.Form.Get("user")

		connect := app.NewConnect()

		var res []*models.Comment

		rows, err := connect.Mysql.Queryx(storage.GetReportComments, movieId, userFio)
		if err != nil {
			utils.ThrowError(err, w)
		} else {
			for rows.Next() {
				var commentReport = models.Comment{}
				err := rows.StructScan(&commentReport)
				if err != nil {
					utils.ThrowError(err, w)
				}
				res = append(res, &commentReport)
			}
		}
		app.CloseConnect(connect)

		json.NewEncoder(w).Encode(res)
		w.WriteHeader(http.StatusOK)
	}
}
