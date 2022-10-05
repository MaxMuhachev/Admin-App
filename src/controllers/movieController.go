package controllers

import (
	"content/src/app"
	"content/src/config"
	"content/src/models"
	"content/src/storage"
	"content/src/utils"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strings"
)

func HandlerMovies(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		RenderMovies(w, nil)
	}
}

func HandlerApiMovies(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var (
		movies        = []models.DataSelect{}
		err           error
		totalElements uint
	)
	search := r.Form.Get("search")
	page := uint(utils.ConvertUint(r.Form.Get("page")))
	movies, err = GetMovieSelect(search, models.DEFAULT_SELECT_LIMIT, page)
	if err == nil {
		totalElements, err = GetSingleResult(storage.GetMoviesCount, search)
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

func HandlerMoviesSearch(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieUserWriteHeader(w, r) {
		res, err := GetMovies(storage.GetMovies)

		app.RenderTemplate(w, "/user/movies/content-movies-search", &app.Page{Title: utils.MOVIE, MovieList: res}, &err)
	}
}

func HandlerPostMoviesFilterApi(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieUserWriteHeader(w, r) {
		r.ParseForm()
		movieString := r.Form.Get("movie")
		genreString := r.Form.Get("genre")

		res, err := GetMoviesByFilter(movieString, genreString)
		if err == nil {
			json.NewEncoder(w).Encode(&res)
			w.WriteHeader(http.StatusOK)
		}
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
		}
	}
}

func GetMoviesByFilter(movie, genre string) ([]*models.Movie, error) {
	connect := app.NewConnect()

	var res []*models.Movie

	rows, err := connect.Mysql.Queryx(storage.GetMoviesByFilter, movie, genre)
	if err != nil {
		return nil, err
	} else {
		for rows.Next() {
			var movie = models.Movie{}
			err := rows.StructScan(&movie)
			if err != nil {
				return nil, err
			}
			res = append(res, &movie)
		}
	}
	app.CloseConnect(connect)
	return res, err
}

func RenderMovies(w http.ResponseWriter, error *error) {
	res, err := GetMovies(storage.GetMovies)

	app.RenderTemplate(w, "movies/content-movies", &app.Page{Title: utils.MOVIE, MovieList: res, Error: error}, &err)
}

func GetMovies(query string) ([]*models.Movie, error) {
	connect := app.NewConnect()

	var res []*models.Movie

	rows, err := connect.Mysql.Queryx(query)
	if err != nil {
		return nil, err
	} else {
		for rows.Next() {
			var movie = models.Movie{}
			err := rows.StructScan(&movie)
			if err != nil {
				return nil, err
			}
			res = append(res, &movie)
		}
	}
	app.CloseConnect(connect)
	return res, err
}

func GetMovieSelect(search string, limit uint, page uint) ([]models.DataSelect, error) {
	connect := app.NewConnect()

	var res []models.DataSelect

	rows, err := connect.Mysql.Queryx(storage.GetMoviesSelect, search, limit, page*limit)
	if err != nil {
		return nil, err
	} else {
		for rows.Next() {
			var movie = models.DataSelect{}
			err := rows.StructScan(&movie)
			if err != nil {
				return nil, err
			}
			res = append(res, movie)
		}
	}
	app.CloseConnect(connect)
	return res, err
}

func GetSingleResult(query string, search string) (uint, error) {
	connect := app.NewConnect()

	var res uint

	err := connect.Mysql.Get(&res, query, search)
	if err != nil {
		return 0, err
	}
	app.CloseConnect(connect)
	return res, err
}

func HandlerCreateMovie(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		genres, err := GetGenres()
		if !utils.ThrowError(err, w) {
			app.RenderTemplate(w, "movies/edit/content-edit-movies", &app.Page{Title: utils.MOVIE, GenreList: genres}, nil)

		}
	}
}

func HandlerCreatePostMovie(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		movie := ParseFormMovie(r)
		if movie.ID == 0 {
			permission, _ := r.Cookie(config.MANAGER_PERSMISSION)
			movie.AddEmpl = permission.Value

			connect := app.NewConnect()
			_, err := connect.Mysql.Queryx(
				storage.CreateMovie,
				movie.Title,
				movie.Year,
				movie.Description,
				movie.KpRating,
				movie.Available,
				movie.VideoLink,
				movie.PictureLink,
				movie.CountView,
				movie.AddEmpl,
			)

			movieRow := connect.Mysql.QueryRow(storage.GetLastMovieId)
			var movieId uint8
			err = movieRow.Scan(&movieId)
			r.Form.Add("id", utils.ConvertToString(movieId))
			movie.ID = movieId

			movieHasGenresMap := ParseMovieGenreForm(r)
			if len(movieHasGenresMap) > 0 && !utils.ThrowError(err, w) {
				if !utils.ThrowError(err, w) {
					genreForDelete, genreForAdd := getMovieHasGenreForChange(make(map[models.MovieLight][]*models.GenreLight), movieHasGenresMap)
					saveMovieHasGenres(w, genreForAdd, genreForDelete)
				}
			}

			genres, err := GetGenres()
			if !utils.ThrowError(err, w) {
				app.RenderTemplate(w, "movies/edit/content-edit-movies", &app.Page{Title: utils.MOVIE, Movie: movie, MovieHasGenreList: movieHasGenresMap, GenreList: genres, Success: utils.MOVIE_SAVED}, &err)
			}
			app.CloseConnect(connect)
		} else {
			HandlerEditPostMovie(w, r)
		}
	}
}

func HandlerEditMovie(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {

		r.ParseForm()
		movieId := r.Form.Get("id")
		var genres []*models.Genre

		connect := app.NewConnect()

		movie, err := GetMovieById(utils.ConvertUint(movieId))
		if !utils.ThrowError(err, w) {
			movieHasGenre, err := GetMovieHasGenres(storage.GetMovieGenreByMovieID, movieId)
			if !utils.ThrowError(err, w) {
				genres, err = GetGenres()
				if !utils.ThrowError(err, w) {
					app.RenderTemplate(w, "movies/edit/content-edit-movies", &app.Page{Title: utils.MOVIE, MovieHasGenreList: movieHasGenre, GenreList: genres, Movie: movie}, &err)
				}
			}
		}
		app.CloseConnect(connect)
	}
}

func HandlerEditPostMovie(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {

		movie := ParseFormMovie(r)
		movieHasGenresMap := ParseMovieGenreForm(r)

		connect := app.NewConnect()

		if len(movieHasGenresMap) > 0 {
			movieHasGenreDB, err := GetMovieHasGenres(storage.GetMovieGenreByMovieID, utils.ConvertToString(movie.ID))
			if !utils.ThrowError(err, w) {
				genreForDelete, genreForAdd := getMovieHasGenreForChange(movieHasGenreDB, movieHasGenresMap)
				saveMovieHasGenres(w, genreForAdd, genreForDelete)
			}
		}

		_, err := connect.Mysql.Queryx(
			storage.UpdateMovie,
			movie.Title,
			movie.Year,
			movie.Description,
			movie.KpRating,
			movie.Available,
			movie.VideoLink,
			movie.PictureLink,
			movie.CountView,
			movie.DateLastEdit,
			movie.ID,
		)

		genres, err := GetGenres()
		if !utils.ThrowError(err, w) {
			app.RenderTemplate(w, "movies/edit/content-edit-movies", &app.Page{Title: utils.MOVIE, Movie: movie, MovieHasGenreList: movieHasGenresMap, GenreList: genres, Success: utils.MOVIE_SAVED}, &err)
		}
		app.CloseConnect(connect)
	}
}

func HandlerDeletePostMovie(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {

		movie := ParseFormMovie(r)
		movie.DateAdd = utils.GetNowString()
		movie.DateLastEdit = movie.DateAdd

		connect := app.NewConnect()
		_, err := connect.Mysql.Queryx(
			storage.DeleteMovieByMovieID,
			movie.ID,
		)
		if err != nil {
			utils.ThrowError(err, w)
		}

		app.CloseConnect(connect)
		w.Write([]byte("1"))
		w.WriteHeader(http.StatusOK)
	}
}

func ParseFormMovie(r *http.Request) *models.Movie {
	r.ParseForm()
	id := r.Form.Get("id")
	title := r.Form.Get("title")
	year := r.Form.Get("year")
	kpRating := r.Form.Get("kpRating")
	available := r.Form.Get("available")
	videoLink := r.Form.Get("videoLink")
	pictureLink := r.Form.Get("pictureLink")
	description := r.Form.Get("description")
	countView := r.Form.Get("countView")
	dateLastEdit := r.Form.Get("dateLastEdit")
	dateAdd := r.Form.Get("dateAdd")
	addEmpl := r.Form.Get("addEmpl")

	var movie = models.Movie{}
	movie.ID = uint8(utils.ConvertUint(id))
	movie.Title = title
	movie.Year = uint32(utils.ConvertUint(year))
	var kpRatingFloat = float32(utils.ConvertFloat(kpRating))
	if kpRatingFloat != 0 {
		movie.KpRating = &kpRatingFloat
	}
	movie.Available = uint(utils.ConvertUint(available))
	movie.VideoLink = videoLink
	movie.PictureLink = pictureLink
	movie.Description = description
	movie.CountView = uint8(utils.ConvertUint(countView))
	movie.DateLastEdit = dateLastEdit
	movie.DateAdd = dateAdd
	movie.AddEmpl = addEmpl
	return &movie
}

func ParseMovieGenreForm(r *http.Request) map[models.MovieLight][]*models.GenreLight {
	r.ParseForm()
	movieId := r.Form.Get("id")
	genreIds := r.Form["genreId[]"]
	var (
		result     = make(map[models.MovieLight][]*models.GenreLight)
		movieLight = models.MovieLight{}
	)
	movieLight.MovieID = uint8(utils.ConvertUint(movieId))

	if len(genreIds) > 0 {
		for _, genreId := range genreIds {
			result[movieLight] = append(result[movieLight], &models.GenreLight{GenreID: uint8(utils.ConvertUint(genreId))})
		}
	} else {
		result[movieLight] = append(result[movieLight], &models.GenreLight{GenreID: 0})
	}

	return result
}

func GetMovieById(id uint64) (*models.Movie, error) {
	connect := app.NewConnect()

	var res *models.Movie

	rows, err := connect.Mysql.Queryx(storage.GetMovieByID, id)
	if err != nil {
		return nil, err
	} else {
		for rows.Next() {
			var movie = models.Movie{}
			err = rows.StructScan(&movie)
			if err != nil {
				return nil, err
			}
			movie.DateLastEdit = strings.Split(movie.DateLastEdit, " ")[0]
			movie.DateAdd = strings.Split(movie.DateAdd, " ")[0]
			res = &movie
		}
	}
	return res, nil
}
