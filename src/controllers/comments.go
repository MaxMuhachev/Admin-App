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

		var res []*models.Comment

		err := app.Conn.Mysql.Select(&res, storage.GetCommentById, userEmail.Value, commentId)
		if !utils.ThrowError(err, w) {
			json.NewEncoder(w).Encode(&res)
			w.WriteHeader(http.StatusOK)
		}

		w.WriteHeader(http.StatusOK)
	}
}

func HandlerGetCommentsByMovie(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieUserWriteHeader(w, r) {
		r.ParseForm()
		movieId := r.Form.Get("movie")
		userEmail, _ := r.Cookie(config.USER_PERSMISSION)

		var res []*models.Comment

		err := app.Conn.Mysql.Select(&res, storage.GetCommentsByMovie, userEmail.Value, movieId)
		if !utils.ThrowError(err, w) {
			json.NewEncoder(w).Encode(&res)
			w.WriteHeader(http.StatusOK)
		}
	}
}

func HandleCreateComment(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieUserWriteHeader(w, r) {
		r.ParseForm()
		movieId := r.Form.Get("movie")
		commentText := r.Form.Get("comment")
		userEmail, _ := r.Cookie(config.USER_PERSMISSION)

		_, err := app.Conn.Mysql.Query(storage.CreateComment, userEmail.Value, movieId, commentText)

		res := models.Comment{}

		err = app.Conn.Mysql.Get(res, storage.GetLastCommentByMovieAndEmail, movieId, userEmail.Value)
		if !utils.ThrowError(err, w) {
			json.NewEncoder(w).Encode(&res)
			w.WriteHeader(http.StatusOK)
		}
	}
}

func HandleUpdatePostComment(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieUserWriteHeader(w, r) {
		r.ParseForm()
		commentId := r.Form.Get("commentId")
		commentText := r.Form.Get("commentText")

		_, err := app.Conn.Mysql.Query(storage.UpdateComment, commentText, commentId)

		res := models.Comment{}

		if !utils.ThrowError(err, w) {
			res.CommentText = commentText
			json.NewEncoder(w).Encode(&res)
			w.WriteHeader(http.StatusOK)
		}
	}
}

func HandleDeleteComment(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieUserWriteHeader(w, r) {
		r.ParseForm()
		id := r.Form.Get("id")

		_, err := app.Conn.Mysql.Query(storage.DeleteComment, id)

		if !utils.ThrowError(err, w) {
			json.NewEncoder(w).Encode("")
			w.WriteHeader(http.StatusOK)
		}
	}
}
