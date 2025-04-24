package repositories

import (
	"backendtickitz/internal/models"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MovieRepository struct {
	db *pgxpool.Pool
}

func NewMovieRepository(db *pgxpool.Pool) *MovieRepository {
	return &MovieRepository{db: db}
}

func (m *MovieRepository) GetMovies(c context.Context) ([]models.MovieStruct, error) {
	query := "SELECT m.id,m.name,  m.duration, m.synopsis  , m.img_movie,m.backdrop, m.release_date, array_agg(g.name) AS genre FROM movie_genre mg JOIN movie m ON mg.movie_id = m.id JOIN genre g ON mg.genre_id = g.id GROUP BY m.id"
	rows, err := m.db.Query(c, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []models.MovieStruct
	for rows.Next() {
		var movie models.MovieStruct
		if err := rows.Scan(&movie.Id, &movie.Name, &movie.Duration, &movie.Synopsis, &movie.Img_movie, &movie.Backdrop, &movie.Release_Date, &movie.Genre); err != nil {

			// c.JSON(http.StatusInternalServerError, gin.H{
			// 	"msg": "terjadi kesalahan sistem",
			// })s
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
