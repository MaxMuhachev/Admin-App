package controllers

import (
	"content/src/app"
	"content/src/models"
	"content/src/storage"
	"content/src/utils"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func HandlerMovieHasGenres(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		res, err := GetMovieHasGenres(storage.GetMovieHasGenre, "")
		app.RenderTemplate(w, "movie-genres/content-movie-genres", &app.Page{Title: utils.MOVIE_HAS_GENRE, MovieHasGenreList: res}, &err)
	}
}

func HandlerCreateMovieHasGenre(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		genres, err := GetGenres()
		if err != nil {
			utils.ThrowError(err, w)
		}
		movies, err := GetMovies(storage.GetMovieWithoutLinkGenre)
		app.RenderTemplate(w, "movie-genres/edit/content-edit-movie-has-genres", &app.Page{Title: utils.MOVIE_HAS_GENRE, GenreList: genres, MovieList: movies}, &err)
	}
}

func HandlerEditMovieHasGenre(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		LoadRenderMovieHasGenre(w, r, "")
	}
}

func LoadRenderMovieHasGenre(w http.ResponseWriter, r *http.Request, successMessageForSave string) {
	r.ParseForm()
	movieId := r.Form.Get("movieId")

	movieHasGenre, err := GetMovieHasGenres(storage.GetMovieGenreByMovieID, movieId)
	var (
		genres []*models.Genre
		movies []*models.Movie
	)

	if err != nil {
		utils.ThrowError(err, w)
	} else {
		genres, err = GetGenres()
		if err != nil {
			utils.ThrowError(err, w)
		} else {
			movies, err = GetMovies(storage.GetMovies)
			if err != nil {
				utils.ThrowError(err, w)
			}
		}
	}

	if successMessageForSave != "" {
		app.RenderTemplate(w, "movie-genres/edit/content-edit-movie-has-genres", &app.Page{Title: utils.GENRES, MovieHasGenreList: movieHasGenre, GenreList: genres, MovieList: movies, Success: successMessageForSave}, &err)
	} else {
		app.RenderTemplate(w, "movie-genres/edit/content-edit-movie-has-genres", &app.Page{Title: utils.GENRES, MovieHasGenreList: movieHasGenre, GenreList: genres, MovieList: movies}, &err)
	}
}

func HandlerEditPostMovieHasGenre(w http.ResponseWriter, r *http.Request) {
	if utils.HasCookieAdminWriteHeader(w, r) {
		movieHasGenresForm := ParseMovieHasGenreForm(r)
		var movieKey models.MovieLight
		for movie, _ := range movieHasGenresForm {
			movieKey = movie
		}
		if len(movieHasGenresForm) > 0 {
			movieHasGenreDB, err := GetMovieHasGenres(storage.GetMovieGenreByMovieID, utils.ConvertToString(movieKey.MovieID))
			if err != nil {
				utils.ThrowError(err, w)
			}
			genreForDelete, genreForAdd := getMovieHasGenreForChange(movieHasGenreDB, movieHasGenresForm)

			saveMovieHasGenres(w, genreForAdd, genreForDelete)
		}
		LoadRenderMovieHasGenre(w, r, utils.MOVIE_HAS_GENRE_SAVED)
	}
}

func saveMovieHasGenres(
	w http.ResponseWriter,
	movieHasGenreForAdd []*models.MovieHasGenre,
	movieHasGenreForDelete []*models.MovieHasGenre,
) {

	for _, movieHasGenre := range movieHasGenreForAdd {
		_, err := app.Conn.Mysql.Queryx(
			storage.CreateMovieHasGenres,
			movieHasGenre.MovieID,
			movieHasGenre.GenreID,
		)
		if err != nil {
			utils.ThrowError(err, w)
		}
	}

	for _, movieHasGenre := range movieHasGenreForDelete {
		_, err := app.Conn.Mysql.Queryx(
			storage.DeleteMovieGenreByMovieGenreID,
			movieHasGenre.MovieID,
			movieHasGenre.GenreID,
		)
		if err != nil {
			utils.ThrowError(err, w)
		}
	}

}

func getMovieHasGenreForChange(
	movieHasGenreDB map[models.MovieLight][]*models.GenreLight,
	movieHasGenresForm map[models.MovieLight][]*models.GenreLight,
) ([]*models.MovieHasGenre, []*models.MovieHasGenre) {
	var (
		genreForAdd    []*models.MovieHasGenre
		genreForDelete []*models.MovieHasGenre
	)
	for movie, genres := range movieHasGenreDB {
		for _, genre := range genres {
			if !utils.Contains(movieHasGenresForm, genre.GenreID) {
				movieHasGenre := models.MovieHasGenre{GenreID: genre.GenreID, MovieID: movie.MovieID}
				genreForDelete = append(genreForDelete, &movieHasGenre)
			}
		}
	}

	for movie, genres := range movieHasGenresForm {
		for _, genre := range genres {
			if !utils.Contains(movieHasGenreDB, genre.GenreID) {
				movieHasGenre := models.MovieHasGenre{GenreID: genre.GenreID, MovieID: movie.MovieID}
				genreForAdd = append(genreForAdd, &movieHasGenre)
			}
		}
	}
	return genreForDelete, genreForAdd
}

func ParseMovieHasGenreForm(r *http.Request) map[models.MovieLight][]*models.GenreLight {
	r.ParseForm()
	movieId := r.Form.Get("movieId")
	genreids := r.Form["genreId[]"]
	var (
		result     = make(map[models.MovieLight][]*models.GenreLight)
		movieLight = models.MovieLight{}
	)
	movieLight.MovieID = uint8(utils.ConvertUint(movieId))

	if len(genreids) > 0 {
		for _, genreId := range genreids {
			result[movieLight] = append(result[movieLight], &models.GenreLight{GenreID: uint8(utils.ConvertUint(genreId))})
		}
	} else {
		result[movieLight] = append(result[movieLight], &models.GenreLight{GenreID: 0})
	}

	return result
}

func GetMovieHasGenres(query string, arg1 string) (map[models.MovieLight][]*models.GenreLight, error) {
	var (
		res  = make(map[models.MovieLight][]*models.GenreLight)
		rows *sqlx.Rows
		err  error
	)

	if arg1 == "" {
		rows, err = app.Conn.Mysql.Queryx(query)
	} else {
		rows, err = app.Conn.Mysql.Queryx(query, arg1)
	}
	if err != nil {
		return nil, err
	} else {
		for rows.Next() {
			var (
				movieHasGenre = models.MovieHasGenre{}
				movieLight    = models.MovieLight{}
				genreLight    = models.GenreLight{}
			)
			err := rows.StructScan(&movieHasGenre)
			if err != nil {
				return nil, err
			}
			movieLight.MovieID = movieHasGenre.MovieID
			movieLight.MovieTitle = movieHasGenre.MovieTitle
			genreLight.GenreID = movieHasGenre.GenreID
			genreLight.GenreTitle = movieHasGenre.GenreTitle

			res[movieLight] = append(res[movieLight], &genreLight)
		}
	}

	return res, nil
}
