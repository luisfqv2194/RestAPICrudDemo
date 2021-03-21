package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"restAPICRUD/movie"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

type Env struct {
	movies movie.MovieService
}

func initDatabase() *Env {
	DBconn, err := sql.Open("sqlite3", "./movie_catalog.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	env := &Env{
		movies: movie.MovieService{Repository: movie.MovieRepository{DB: DBconn}},
	}

	//Create tables if not exists
	env.movies.Repository.CreateTable()

	return env
}

func setupRoutes(env *Env) http.Handler {

	router := httprouter.New()
	apiV1 := "/api/v1"
	// movie Handlers
	router = movie.MakeHandlers(apiV1, router, env.movies)

	return router
}

const defaultPort = "8080"

func main() {
	// Set configuration variables
	env := initDatabase()
	var (
		addr     = envString("PORT", defaultPort)
		httpAddr = flag.String("http.addr", ":"+addr, "HTTP listen address")
	)
	flag.Parse()
	// Create a logger
	logger := log.New(os.Stdout, "", log.Lshortfile)
	if err := http.ListenAndServe(*httpAddr, setupRoutes(env)); err != nil {
		logger.Fatal(err)
	}
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
