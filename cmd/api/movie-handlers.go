package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.Logger.Println(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	app.Logger.Println("id is", id)

	movie, err := app.Models.DB.Get(id)

	if err != nil {
		app.Logger.Println(err)
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, movie, "movie")
	if err != nil {
		app.Logger.Println(err)
		app.errorJSON(w, err)
		return
	}

}

func (app *Application) getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.Models.DB.All()
	if err != nil {
		app.Logger.Println(err)
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		app.Logger.Println(err)
		app.errorJSON(w, err)
		return
	}
}
func (app *Application) getAllGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := app.Models.DB.GenresAll()
	if err != nil {
		app.Logger.Println(err)
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, genres, "genres")
	if err != nil {
		app.Logger.Println(err)
		app.errorJSON(w, err)
		return
	}
}
func (app *Application) getAllMoviesByGenre(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	genreID, err := strconv.Atoi(params.ByName("genre_id"))
	if err != nil {
		app.Logger.Println(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	movies, err := app.Models.DB.All(genreID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		app.Logger.Println(err)
		app.errorJSON(w, err)
		return
	}
}

func (app *Application) deleteMovie(w http.ResponseWriter, r *http.Request) {

}

func (app *Application) insertMovie(w http.ResponseWriter, r *http.Request) {

}

func (app *Application) updateMovie(w http.ResponseWriter, r *http.Request) {

}

func (app *Application) searchMovie(w http.ResponseWriter, r *http.Request) {

}
