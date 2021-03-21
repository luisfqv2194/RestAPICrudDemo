package movie

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type MovieService struct {
	Repository MovieRepository
}

func MakeHandlers(apiVersion string, router *httprouter.Router, movieService MovieService) *httprouter.Router {
	router.POST(apiVersion+"/movies", movieService.createMovieHandler())
	router.GET(apiVersion+"/movies", movieService.listMoviesHandler())
	router.GET(apiVersion+"/movies/:id", movieService.getMovieHandler())
	router.PUT(apiVersion+"/movies/:id", movieService.updateMovieHandler())
	router.DELETE(apiVersion+"/movies/:id", movieService.deleteMovieHandler())
	return router
}

func (m MovieService) createMovieHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		decoder := json.NewDecoder(r.Body)
		decoder.UseNumber()
		var movie Movie
		decoder.Decode(&movie)
		if err := m.Repository.CreateMovie(movie); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(movie)

	}
}

func (m MovieService) listMoviesHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		movieList, err := m.Repository.FindAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(movieList)

	}
}

func (m MovieService) getMovieHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		_, err := strconv.Atoi(params.ByName("id"))
		if err != nil {
			http.Error(w, "id is invalid", http.StatusBadRequest)
			return
		}
		movie, err := m.Repository.Find(params.ByName("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(movie)

	}
}

func (m MovieService) updateMovieHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		_, err := strconv.Atoi(params.ByName("id"))
		if err != nil {
			http.Error(w, "id is invalid", http.StatusBadRequest)
			return
		}
		decoder := json.NewDecoder(r.Body)
		decoder.UseNumber()
		var movie Movie
		decoder.Decode(&movie)
		if err := m.Repository.Update(params.ByName("id"), movie); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(movie)

	}
}

func (m MovieService) deleteMovieHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		_, err := strconv.Atoi(params.ByName("id"))
		if err != nil {
			http.Error(w, "id is invalid", http.StatusBadRequest)
			return
		}
		if err := m.Repository.Remove(params.ByName("id")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode("Ok")

	}
}
