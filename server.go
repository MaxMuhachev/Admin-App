package main

import (
	"content/src/controllers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("src/resources/static"))))

	addGetRoute("/", router, controllers.HandlerIndex)
	addPostRoute("/", router, controllers.HandlerPostIndex)

	addGetRoute("/api/movies", router, controllers.HandlerApiMovies)
	addGetRoute("/movies", router, controllers.HandlerMovies)
	addGetRoute("/movies/edit", router, controllers.HandlerEditMovie)
	addPostRoute("/movies/edit", router, controllers.HandlerEditPostMovie)
	addGetRoute("/movies/create", router, controllers.HandlerCreateMovie)
	addPostRoute("/movies/create", router, controllers.HandlerCreatePostMovie)
	addPostRoute("/movies/delete", router, controllers.HandlerDeletePostMovie)

	addGetRoute("/genres", router, controllers.HandlerGenres)
	addGetRoute("/genres/edit", router, controllers.HandlerEditGenre)
	addPostRoute("/genres/edit", router, controllers.HandlerEditPostGenre)
	addGetRoute("/genres/create", router, controllers.HandlerCreateGenre)
	addPostRoute("/genres/create", router, controllers.HandlerCreatePostGenre)
	addPostRoute("/genres/delete", router, controllers.HandlerDeletePostGenre)

	addGetRoute("/comments", router, controllers.HandlerViewComments)

	addGetRoute("/statistic", router, controllers.HandlerStatistic)
	addGetRoute("/statistic/films", router, controllers.HandlerStatisticFilm)
	addGetRoute("/statistic/films/get", router, controllers.HandlerStatisticFilmWithDate)
	addGetRoute("/statistic/users", router, controllers.HandlerStatisticUsers)
	addGetRoute("/statistic/users/get", router, controllers.HandlerStatisticUsersWithDate)
	addGetRoute("/statistic/comments", router, controllers.HandlerStatisticComments)
	addGetRoute("/statistic/comments/get", router, controllers.HandlerStatisticCommentsWithDate)

	addGetRoute("/reports", router, controllers.HandlerReport)
	addGetRoute("/reports/films", router, controllers.HandlerReportFilm)
	addGetRoute("/reports/films/get", router, controllers.HandlerReportFilmWithDate)
	addGetRoute("/reports/users", router, controllers.HandlerReportUsers)
	addGetRoute("/reports/users/get", router, controllers.HandlerReportUsersWithDate)
	addGetRoute("/reports/comments", router, controllers.HandlerReportComments)
	addGetRoute("/reports/comments/get", router, controllers.HandlerReportCommentsByMovieUser)

	addGetRoute("/user/edit", router, controllers.HandlerEditUser)
	addPostRoute("/user/edit", router, controllers.HandlerEditPostUser)
	addGetRoute("/api/users", router, controllers.HandlerApiUsers)

	addGetRoute("/api/company", router, controllers.HandlerApiCompany)
	addGetRoute("/employees", router, controllers.HandlerEmployees)
	addGetRoute("/employees/edit", router, controllers.HandlerEditEmployee)
	addPostRoute("/employees/edit", router, controllers.HandlerEditPostEmployee)
	addGetRoute("/employees/create", router, controllers.HandlerCreateEmployee)
	addPostRoute("/employees/create", router, controllers.HandlerCreatePostEmployee)
	addPostRoute("/employees/delete", router, controllers.HandlerDeletePostEmployee)

	addGetRoute("/movies-search", router, controllers.HandlerMoviesSearch)
	addGetRoute("/movies-filter", router, controllers.HandlerPostMoviesFilterApi)
	addGetRoute("/user/comments", router, controllers.HandlerUserComments)
	addGetRoute("/user/comment/get", router, controllers.HandlerGetCommentById)
	addGetRoute("/user/comments/get", router, controllers.HandlerGetCommentsByMovie)
	addPostRoute("/user/comment/create", router, controllers.HandleCreateComment)
	addPostRoute("/user/comment", router, controllers.HandleUpdatePostComment)
	addPostRoute("/user/comment/delete", router, controllers.HandleDeleteComment)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}

func addGetRoute(path string, router *mux.Router, handler http.HandlerFunc) {
	router.HandleFunc(path, handler).Methods("GET")
}

func addPostRoute(path string, router *mux.Router, handler http.HandlerFunc) {
	router.HandleFunc(path, handler).Methods("POST")
}
