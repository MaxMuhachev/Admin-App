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

func HandlerUserComments(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieUserWriteHeader(w, r) {
		app.RenderTemplate(w, "user/comments/content-comments", &app.Page{}, nil)
	}
}

func HandlerViewComments(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieUserWriteHeader(w, r) {
		app.RenderTemplate(w, "comments/content-comments", &app.Page{}, nil)
	}
}

func HandlerGetCommentById(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieUserWriteHeader(w, r) {
		r.ParseForm()
		commentId := r.Form.Get("commentId")
		userEmail, _ := r.Cookie(config.USER_PERSMISSION)

		connect := app.NewConnect()

		var res []*models.Comment

		rows, err := connect.Mysql.Queryx(storage.GetCommentById, userEmail.Value, commentId)
		if !utils.ThrowError(err, w) {
			for rows.Next() {
				var comment = models.Comment{}
				err := rows.StructScan(&comment)
				if err != nil {
					utils.ThrowError(err, w)
				}
				res = append(res, &comment)
			}
		}
		app.CloseConnect(connect)

		json.NewEncoder(w).Encode(res)
		w.WriteHeader(http.StatusOK)
	}
}

func HandlerGetCommentsByMovie(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieUserWriteHeader(w, r) {
		r.ParseForm()
		movieId := r.Form.Get("movie")
		userEmail, _ := r.Cookie(config.USER_PERSMISSION)

		connect := app.NewConnect()

		var res []*models.Comment

		rows, err := connect.Mysql.Queryx(storage.GetCommentsByMovie, userEmail.Value, movieId)
		if err != nil {
			utils.ThrowError(err, w)
		} else {
			for rows.Next() {
				var comment = models.Comment{}
				err := rows.StructScan(&comment)
				if err != nil {
					utils.ThrowError(err, w)
				}
				res = append(res, &comment)
			}
		}
		app.CloseConnect(connect)

		json.NewEncoder(w).Encode(res)
		w.WriteHeader(http.StatusOK)
	}
}

func HandleCreateComment(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieUserWriteHeader(w, r) {
		r.ParseForm()
		movieId := r.Form.Get("movie")
		commentText := r.Form.Get("comment")
		userEmail, _ := r.Cookie(config.USER_PERSMISSION)

		connect := app.NewConnect()

		_, err := connect.Mysql.Queryx(storage.CreateComment, userEmail.Value, movieId, commentText)

		res := models.Comment{}

		rows, err := connect.Mysql.Queryx(storage.GetLastCommentByMovieAndEmail, movieId, userEmail.Value)
		if err != nil {
			utils.ThrowError(err, w)
		} else {
			for rows.Next() {
				err := rows.StructScan(&res)
				if err != nil {
					utils.ThrowError(err, w)
				} else {
					json.NewEncoder(w).Encode(&res)
					w.WriteHeader(http.StatusOK)
				}
			}
		}
		app.CloseConnect(connect)
	}
}

func HandleUpdatePostComment(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieUserWriteHeader(w, r) {
		r.ParseForm()
		commentId := r.Form.Get("commentId")
		commentText := r.Form.Get("commentText")

		connect := app.NewConnect()

		_, err := connect.Mysql.Queryx(storage.UpdateComment, commentText, commentId)

		res := models.Comment{}

		if !utils.ThrowError(err, w) {
			res.CommentText = commentText
			json.NewEncoder(w).Encode(&res)
			w.WriteHeader(http.StatusOK)
		}
		app.CloseConnect(connect)
	}
}

func HandleDeleteComment(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieUserWriteHeader(w, r) {
		r.ParseForm()
		id := r.Form.Get("id")
		connect := app.NewConnect()

		_, err := connect.Mysql.Queryx(storage.DeleteComment, id)

		if err != nil {
			utils.ThrowError(err, w)
		} else {
			json.NewEncoder(w).Encode("")
		}
		app.CloseConnect(connect)
	}
}
