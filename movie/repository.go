package movie

import (
	"database/sql"
	"log"
)

type Movie struct {
	ID       uint
	Rating   string
	Title    string
	Duration int
	Year     int
}

type MovieRepository struct {
	DB *sql.DB
}

func (m MovieRepository) CreateMovie(movie Movie) error {
	log.Println("Inserting movie record ...")
	insertMovieSQL := `INSERT INTO movie(rating, title, duration, year) VALUES (?, ?, ?, ?)`
	statement, err := m.DB.Prepare(insertMovieSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Println(err.Error())
		return err

	}
	_, err = statement.Exec(movie.Rating, movie.Title, movie.Duration, movie.Year)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return err
}

func (m MovieRepository) FindAll() ([]Movie, error) {
	log.Println("Getting all movies.")
	movieList := []Movie{}
	getMoviesSQL := `Select * from movie`
	rows, err := m.DB.Query(getMoviesSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Println(err.Error())
		return movieList, err
	}

	for rows.Next() {
		var movieResult Movie
		if err := rows.Scan(&movieResult.ID, &movieResult.Rating, &movieResult.Title, &movieResult.Duration, &movieResult.Year); err != nil {
			log.Println(err.Error())
			return movieList, err
		}
		movieList = append(movieList, movieResult)

	}

	return movieList, err
}

func (m MovieRepository) Find(id string) (Movie, error) {
	log.Println("Getting movie.")
	getMovieSQL := "Select * from movie where id =" + id
	row := m.DB.QueryRow(getMovieSQL) // Prepare statement.

	var movieResult Movie
	err := row.Scan(&movieResult.ID, &movieResult.Rating, &movieResult.Title, &movieResult.Duration, &movieResult.Year)

	return movieResult, err
}

func (m MovieRepository) CreateTable() error {
	createStudentTableSQL := `CREATE TABLE if not exists movie (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"rating" TEXT,
		"title" TEXT,
		"duration" INTEGER,
		"year" INTEGER		
	  );` // SQL Statement for Create Table

	statement, err := m.DB.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Println(err.Error())
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		log.Println(err.Error())
		return err

	}
	return err
}

func (m MovieRepository) Update(id string, movie Movie) error {
	updateStatement :=
		"UPDATE movie SET rating = ?, title = ?, duration = ?, year = ? WHERE id = ?"
	statement, err := m.DB.Prepare(updateStatement)
	// This is good to avoid SQL injections
	if err != nil {
		log.Println(err.Error())
		return err

	}
	_, err = statement.Exec(movie.Rating, movie.Title, movie.Duration, movie.Year, id)
	if err != nil {
		log.Println(err.Error())
		return err

	}

	return err

}

func (m MovieRepository) Remove(id string) error {
	log.Println("Deleting movie.")
	deleteMovieSQL := "DELETE FROM movie WHERE id = ?"
	statement, err := m.DB.Prepare(deleteMovieSQL)
	// This is good to avoid SQL injections
	if err != nil {
		log.Println(err.Error())
		return err
	}
	_, err = statement.Exec(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return err

}
