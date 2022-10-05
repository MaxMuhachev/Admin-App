package app

import "content/src/models"

type Page struct {
	Title             string
	MovieList         []*models.Movie
	MovieLightList    []*models.MovieLight
	Movie             *models.Movie
	MovieHasGenreList map[models.MovieLight][]*models.GenreLight
	GenreList         []*models.Genre
	Employee          *models.Employee
	EmployeeList      []*models.Employee
	User              *models.User
	Genre             *models.Genre

	MovieStatistic []*models.MovieReport

	Error   *error
	Success string
}
