package repositories

import (
	"backendtickitz/internal/models"
	"backendtickitz/pkg"
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MovieRepository struct {
	db *pgxpool.Pool
}

func NewMovieRepository(db *pgxpool.Pool) *MovieRepository {
	return &MovieRepository{db: db}
}

func (m *MovieRepository) GetMovies(c context.Context, params *models.MovieQueryParams) ([]models.MovieStruct, error) {
	offset := (params.Page - 1) * params.Pagesize
	args := []any{params.Pagesize, offset}
	query := "SELECT m.id,m.name,  m.duration, m.synopsis  , m.img_movie,m.backdrop, m.release_date, array_agg(g.name) AS genre FROM movie_genre mg JOIN movie m ON mg.movie_id = m.id JOIN genre g ON mg.genre_id = g.id GROUP BY m.id LIMIT $1 OFFSET $2 "
	rows, err := m.db.Query(c, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []models.MovieStruct
	for rows.Next() {
		var movie models.MovieStruct
		if err := rows.Scan(&movie.Id, &movie.Name, &movie.Duration, &movie.Synopsis, &movie.Img_movie, &movie.Backdrop, &movie.Release_Date, &movie.Genre); err != nil {
			return nil, err
		}
		result = append(result, movie)
	}
	return result, nil
}

func (m *MovieRepository) GetMovieById(c context.Context, idInt int) (models.MovieStruct, error) {

	query := `SELECT m.id,m.name,  m.duration, m.synopsis  , m.img_movie,m.backdrop, m.release_date, array_agg(g.name) AS genre FROM movie_genre mg JOIN movie m ON mg.movie_id = m.id JOIN genre g ON mg.genre_id = g.id where m.id= $1 GROUP BY m.id`
	var result models.MovieStruct
	if err := m.db.QueryRow(c, query, idInt).Scan(&result.Id, &result.Name, &result.Duration, &result.Synopsis, &result.Img_movie, &result.Backdrop, &result.Release_Date, &result.Genre); err != nil {
		return models.MovieStruct{}, err
	}
	return result, nil
}

func (m *MovieRepository) AddMovie(c context.Context, newMovie *models.MovieStruct) error {
	query := "INSERT INTO movie (name, duration, synopsis, img_movie, backdrop, release_date) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	queryAddGenre := "INSERT INTO movie_genre (movie_id, genre_id) VALUES ($1, $2)"
	log.Println("Executing query", query)
	var id int
	value := []any{newMovie.Name, newMovie.Duration, newMovie.Synopsis, newMovie.Img_movie, newMovie.Backdrop, newMovie.Release_Date}
	err := pkg.DB.QueryRow(c, query, value...).Scan(&id)
	log.Printf("Movie created with ID: %v", id)
	log.Println("ini oi oi oi")
	for _, genreId := range newMovie.Genre {
		log.Println("ini id genre", genreId)
		_, err := pkg.DB.Exec(c, queryAddGenre, id, genreId)
		if err != nil {
			log.Println("ini error", err)
			return err
		}
	}

	if err != nil {
		return err
	}
	return nil
}

func (m *MovieRepository) UpdateMovie(c context.Context, newMovie *models.MovieStruct, idMovie int) error {
	query := "update movie set name =$1, duration = $2, synopsis = $3, img_movie = $4, backdrop=$5 where id = $6"
	value := []any{newMovie.Name, newMovie.Duration, newMovie.Synopsis, newMovie.Img_movie, newMovie.Backdrop, idMovie}
	cmd, err := pkg.DB.Exec(c, query, value...)
	if err != nil {
		return err
	}
	if rowsAffected := cmd.RowsAffected(); rowsAffected == 0 {
		return errors.New("no rows were updated")
	}
	return nil
}

func (m *MovieRepository) DeleteMovie(c context.Context, idMovie int) (*models.ResultCommand, error) {
	tx, err := pkg.DB.Begin(c)
	if err != nil {
		return nil, err
	}
	log.Println("ini data param", idMovie)
	defer func() {
		if err != nil {
			tx.Rollback(c)
		}
	}()
	queryDeleteAsosiasi := "delete from movie_genre where movie_id = $1"
	cmdMovieGenre, err := tx.Exec(c, queryDeleteAsosiasi, idMovie)
	if err != nil {
		return nil, err
	}
	queryDeleteMovie := "delete from movie where id = $1"
	cmdMovie, err := tx.Exec(c, queryDeleteMovie, idMovie)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(c); err != nil {
		return nil, err
	}
	return &models.ResultCommand{
		ResultFirst:  cmdMovie,
		ResultSecond: cmdMovieGenre,
	}, nil
}
