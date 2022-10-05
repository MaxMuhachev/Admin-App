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

		connect := app.NewConnect()

		var res uint

		rows, err := connect.Mysql.Queryx(storage.GetStatisticMovie, startDate, endDate)
		if err != nil {
			utils.ThrowError(err, w)
		} else {
			for rows.Next() {
				err := rows.Scan(&res)
				if err != nil {
					utils.ThrowError(err, w)
				}
			}
		}

		w.Write([]byte(utils.ConvertToString(uint8(res))))
		w.WriteHeader(http.StatusOK)
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

		connect := app.NewConnect()

		var res uint

		rows, err := connect.Mysql.Queryx(storage.GetStatisticUsers, startDate, endDate)
		if err != nil {
			utils.ThrowError(err, w)
		} else {
			for rows.Next() {
				err := rows.Scan(&res)
				if err != nil {
					utils.ThrowError(err, w)
				}
			}
		}

		w.Write([]byte(utils.ConvertToString(uint8(res))))
		w.WriteHeader(http.StatusOK)
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

		connect := app.NewConnect()

		var res []*models.CommentStatistic

		rows, err := connect.Mysql.Queryx(storage.GetStatisticComments, startDate, endDate)
		if err != nil {
			utils.ThrowError(err, w)
		} else {
			for rows.Next() {
				var movieStatistic = models.CommentStatistic{}
				err := rows.StructScan(&movieStatistic)
				if err != nil {
					utils.ThrowError(err, w)
				}
				res = append(res, &movieStatistic)
			}
		}

		json.NewEncoder(w).Encode(res)
		w.WriteHeader(http.StatusOK)
	}
}
