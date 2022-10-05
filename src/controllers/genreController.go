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
	connect := app.NewConnect()

	var res []*models.Genre

	rows, err := connect.Mysql.Queryx(storage.GetGenres)
	if err != nil {
		return nil, err
	} else {
		for rows.Next() {
			var genre = models.Genre{}
			err := rows.StructScan(&genre)
			if err != nil {
				return nil, err
			}
			res = append(res, &genre)
		}
	}
	app.CloseConnect(connect)
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

		connect := app.NewConnect()
		_, err := connect.Mysql.Queryx(
			storage.CreateGenre,
			genre.Title,
			genre.Description,
		)

		if !utils.ThrowError(err, w) {
			var genreId uint8
			genreRow := connect.Mysql.QueryRow(storage.GetGenreIdLast)
			err = genreRow.Scan(&genreId)
			if !utils.ThrowError(err, w) {
				r.Form.Set("id", utils.ConvertToString(genreId))
				http.Redirect(w, r, "/genres/edit", http.StatusTemporaryRedirect)
			}
		}
		app.CloseConnect(connect)
	}
}

func HandlerEditGenre(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {

		r.ParseForm()
		id := r.Form.Get("id")

		connect := app.NewConnect()

		genre, err := GetGenreById(utils.ConvertUint(id))
		app.RenderTemplate(w, "genres/edit/content-edit-genres", &app.Page{Title: utils.GENRES, Genre: genre}, &err)

		app.CloseConnect(connect)
	}
}

func HandlerEditPostGenre(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		genre := ParseGenreForm(r)
		connect := app.NewConnect()

		_, err := connect.Mysql.Queryx(
			storage.UpdateGenre,
			genre.Title,
			genre.Description,
			genre.ID,
		)

		app.CloseConnect(connect)
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
	connect := app.NewConnect()

	var res *models.Genre

	rows, err := connect.Mysql.Queryx(storage.GetGenreByID, id)
	if err != nil {
		return nil, err
	} else {
		for rows.Next() {
			var genre = models.Genre{}
			err = rows.StructScan(&genre)
			if err != nil {
				return nil, err
			}
			res = &genre
		}
	}
	return res, nil
}

func HandlerDeletePostGenre(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		genre := ParseGenreForm(r)

		connect := app.NewConnect()
		_, err := connect.Mysql.Queryx(
			storage.DeleteGenreByGenreID,
			genre.ID,
		)
		if err != nil {
			utils.ThrowError(err, w)
		}

		app.CloseConnect(connect)
		w.Write([]byte("1"))
		w.WriteHeader(http.StatusOK)
	}
}
