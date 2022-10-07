package controllers

import (
	"content/src/app"
	"content/src/models"
	"content/src/storage"
	"content/src/utils"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func HandlerGenres(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		res, err := GetGenres()
		app.RenderTemplate(w, "genres/content-genres", &app.Page{Title: utils.GENRES, GenreList: res}, &err)
	}
}

func GetGenres() ([]*models.Genre, error) {

	var res []*models.Genre

	err := app.Conn.Mysql.Select(&res, storage.GetGenres)
	if err != nil {
		return nil, err
	}
	return res, err
}

func HandlerCreateGenre(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		app.RenderTemplate(w, "genres/edit/content-edit-genres", &app.Page{Title: utils.GENRES}, nil)
	}
}

func HandlerCreatePostGenre(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {

		genre := ParseGenreForm(r)

		_, err := app.Conn.Mysql.Query(
			storage.CreateGenre,
			genre.Title,
			genre.Description,
		)

		if !utils.ThrowError(err, w) {
			var genreId uint8
			genreRow := app.Conn.Mysql.QueryRow(storage.GetGenreIdLast)
			err = genreRow.Scan(&genreId)
			if !utils.ThrowError(err, w) {
				r.Form.Set("id", utils.ConvertToString(genreId))
				http.Redirect(w, r, "/genres/edit", http.StatusTemporaryRedirect)
			}
		}

	}
}

func HandlerEditGenre(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {

		r.ParseForm()
		id := r.Form.Get("id")

		genre, err := GetGenreById(utils.ConvertUint(id))
		app.RenderTemplate(w, "genres/edit/content-edit-genres", &app.Page{Title: utils.GENRES, Genre: genre}, &err)
	}
}

func HandlerEditPostGenre(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		genre := ParseGenreForm(r)

		_, err := app.Conn.Mysql.Query(
			storage.UpdateGenre,
			genre.Title,
			genre.Description,
			genre.ID,
		)

		app.RenderTemplate(w, "genres/edit/content-edit-genres", &app.Page{Title: utils.GENRES, Genre: genre, Success: utils.GENRE_SAVED}, &err)
	}
}

func ParseGenreForm(r *http.Request) *models.Genre {
	r.ParseForm()
	id := r.Form.Get("id")
	title := r.Form.Get("title")
	description := r.Form.Get("description")

	var movie = models.Genre{}
	movie.ID = uint8(utils.ConvertUint(id))
	movie.Title = title
	movie.Description = description
	return &movie
}

func GetGenreById(id uint64) (*models.Genre, error) {
	res := &models.Genre{}

	err := app.Conn.Mysql.Get(res, storage.GetGenreByID, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func HandlerDeletePostGenre(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		genre := ParseGenreForm(r)

		_, err := app.Conn.Mysql.Query(
			storage.DeleteGenreByGenreID,
			genre.ID,
		)
		if err != nil {
			utils.ThrowError(err, w)
		}

		w.Write([]byte("1"))
		w.WriteHeader(http.StatusOK)
	}
}
