package repositories

import (
	"backendtickitz/internal/models"
	"backendtickitz/pkg"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type MovieRepository struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

func NewMovieRepository(db *pgxpool.Pool, rdb *redis.Client) *MovieRepository {
	return &MovieRepository{db: db, rdb: rdb}
}

// func (m *MovieRepository) GetProduct(c context.Context, params *models.ProductQueryParams) ([]models.ProductStruct, error) {
// 	pagesize := 6
// 	offset := (params.Page - 1) * pagesize
// 	args := []any{pagesize, offset}
// 	paramIndex := 3
// 	whereClauses := []string{}
// 	query := ""

// 	if params.Name != "" {
// 		whereClauses = append(whereClauses, fmt.Sprintf("m.name ILIKE $%d", paramIndex))
// 		args = append(args, "%"+params.Name+"%")
// 		paramIndex++
// 	}
// 	if params.Name != "" {
// 		whereClauses = append(whereClauses, fmt.Sprintf("m.name ILIKE $%d", paramIndex))
// 		args = append(args, "%"+params.Name+"%")
// 		paramIndex++
// 	}

// }

func (m *MovieRepository) GetMovies(c context.Context, params *models.MovieQueryParams) ([]models.MovieStruct, error) {
	pagesize := 12
	offset := (params.Page - 1) * pagesize
	args := []any{pagesize, offset}
	paramIndex := 3
	query := `SELECT m.id,m.name,  m.duration, m.synopsis, m.img_movie,m.backdrop, m.release_date, array_agg(g.name) AS genre FROM movie_genre mg JOIN movie m ON mg.movie_id = m.id JOIN genre g ON mg.genre_id = g.id `
	whereClauses := []string{}
	if params.Name != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("m.name ILIKE $%d", paramIndex))
		args = append(args, "%"+params.Name+"%")
		paramIndex++
	}
	// Gabungkan kondisi WHERE
	if params.Genre != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("g.name = $%d", paramIndex))
		args = append(args, params.Genre)

		paramIndex++
	}
	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	query += " GROUP BY m.id, m.name, m.duration, m.synopsis, m.img_movie, m.backdrop, m.release_date limit $1 offset $2  "

	rows, err := m.db.Query(c, query, args...)
	if err != nil {
		return nil, err
	}
	log.Println("query", query)
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

func (m *MovieRepository) Upcoming(c context.Context) ([]models.MovieStruct, error) {

	redisKey := "movie"
	cache, err := m.rdb.Get(c, redisKey).Result()
	if err != nil {
		if err == redis.Nil {
			log.Printf("\nkey %s does not exist\n", redisKey)
		} else {
			log.Println("Redis not working")
		}
	} else {
		var movie []models.MovieStruct
		if err := json.Unmarshal([]byte(cache), &movie); err != nil {
			return nil, err
		}
		if len(movie) > 0 {
			return movie, nil
		}

	}
	// query := `SELECT m.id as id, m."name" as name,m.duration as duration, m.img_movie as img_movie, m.release_date,json_agg(distinct g."name") as genre FROM orders o  join schedule s on s.id  = o.id_schedule JOIN movie m ON m.id = s.id_movie JOIN movie_genre mg ON mg.movie_id = m.id JOIN genre g ON g.id = mg.genre_id where m.release_date > now() group by m.id LIMIT $1`
	query := `SELECT m.id, m.name as name, m.img_movie , m.release_date,json_agg(g."name") FROM movie m JOIN movie_genre mg ON mg.movie_id = m.id JOIN genre g ON g.id = mg.genre_id where m.release_date > now() group by m.id limit $1 `
	limit := 5
	var result []models.MovieStruct
	rows, err := m.db.Query(c, query, limit)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var movie models.MovieStruct
		if err := rows.Scan(&movie.Id, &movie.Name, &movie.Img_movie, &movie.Release_Date, &movie.Genre); err != nil {
			return nil, err
		}
		result = append(result, movie)
	}

	// data baru yang diambil masukkan ke dalam redis
	res, err := json.Marshal(result)
	if err != nil {
		log.Println("[DEBUG] marshal", err.Error())
	}
	if err := m.rdb.Set(c, redisKey, string(res), time.Minute*5).Err(); err != nil {
		log.Println("DEBUG redos set", err.Error())
	}
	return result, nil
}
func (m *MovieRepository) Popular(c context.Context) ([]models.MovieStruct, error) {
	query := `SELECT m.id ,m.name , m.img_movie,m.release_date,json_agg(distinct g."name") as genre FROM movie m JOIN movie_genre mg ON mg.movie_id = m.id JOIN genre g ON g.id = mg.genre_id LEFT JOIN schedule s ON s.id_movie = m.id LEFT JOIN orders o ON o.id_schedule = s.id LEFT JOIN orders_seats os ON os.orders_id = o.id where m.release_date < NOW() GROUP BY  m.id having  COUNT(DISTINCT os.seats_id) > 1 ORDER BY  COUNT(DISTINCT os.seats_id) DESC LIMIT $1`
	limit := 5
	var result []models.MovieStruct
	rows, err := m.db.Query(c, query, limit)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var movie models.MovieStruct
		if err := rows.Scan(&movie.Id, &movie.Name, &movie.Img_movie, &movie.Release_Date, &movie.Genre); err != nil {
			return nil, err
		}
		result = append(result, movie)
	}
	return result, nil
}

func (m *MovieRepository) GetMovieById(c context.Context, idInt int) (models.MovieStruct, error) {
	log.Println("[ini detail]", idInt)

	query := `SELECT m.id,m.name,  m.duration, m.synopsis, m.img_movie,m.backdrop, m.release_date, array_agg(g.name) AS genre FROM movie_genre mg JOIN movie m ON mg.movie_id = m.id JOIN genre g ON mg.genre_id = g.id where m.id= $1 GROUP BY m.id`
	var result models.MovieStruct

	if err := m.db.QueryRow(c, query, idInt).Scan(&result.Id, &result.Name, &result.Duration, &result.Synopsis, &result.Img_movie, &result.Backdrop, &result.Release_Date, &result.Genre); err != nil {
		log.Println("[DEBUG M DETAIL]", err)
		return models.MovieStruct{}, err
	}
	return result, nil
}

func (m *MovieRepository) AddMovie(c context.Context, newMovie *models.MovieFrom, uploadedFiles map[string]any) error {
	tx, err := m.db.Begin(c)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(c)
		}
	}()
	query := "INSERT INTO movie (name, duration, synopsis, img_movie, backdrop, release_date) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	queryAddGenre := "INSERT INTO movie_genre (movie_id, genre_id) VALUES ($1, $2)"
	log.Println("Executing query", query)
	var id int
	value := []any{newMovie.Name, newMovie.Duration, newMovie.Synopsis, uploadedFiles["img_movie"], uploadedFiles["backdrop"], newMovie.Release_Date}
	err = tx.QueryRow(c, query, value...).Scan(&id)
	for _, genreId := range newMovie.Genre {
		_, err := tx.Exec(c, queryAddGenre, id, genreId)
		if err != nil {
			log.Println("ini error", err)
			return err
		}
	}
	log.Println("[DEBUG ID CINEMA AND PRICE]", newMovie.IdCinema, newMovie.Price)
	queryAddSchedule := "insert into schedule (id_movie, id_cinema, price,date) values ($1,$2,$3,$4)"
	valueSchedule := []any{id, newMovie.IdCinema, newMovie.Price, newMovie.Release_Date}
	cmd, err := tx.Exec(c, queryAddSchedule, valueSchedule...)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		tx.Rollback(c)
		return errors.New("add schedule failed")
	}
	if err := tx.Commit(c); err != nil {
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

func (m *MovieRepository) DeleteMovie(c context.Context, movies *models.MovieStruct) (*models.ResultCommand, error) {
	tx, err := pkg.DB.Begin(c)
	if err != nil {
		return nil, err
	}
	log.Println("ini data param", movies.Id)
	defer func() {
		if err != nil {
			tx.Rollback(c)
		}
	}()
	queryDeleteAsosiasi := "delete from movie_genre where movie_id = $1"
	cmdMovieGenre, err := tx.Exec(c, queryDeleteAsosiasi, movies.Id)
	if err != nil {
		return nil, err
	}
	queryDeleteMovie := "delete from movie where id = $1"
	cmdMovie, err := tx.Exec(c, queryDeleteMovie, movies.Id)
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

func (m *MovieRepository) GetGenres(c context.Context) ([]models.Genre, error) {
	query := `SELECT id, name FROM genre`

	rows, err := m.db.Query(c, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query genres: %w", err)
	}
	defer rows.Close()

	var genres []models.Genre
	for rows.Next() {
		var genre models.Genre
		if err := rows.Scan(&genre.Id, &genre.Name); err != nil {
			return nil, fmt.Errorf("failed to scan genre: %w", err)
		}
		genres = append(genres, genre)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return genres, nil
}

func (m *MovieRepository) GetCinema(c context.Context) ([]models.Cinema, error) {
	query := `select cl.id, c.name as name,l.name as location  from cinema c  join cinema_location cl on cl.cinema_id = c.id join location l on l.id  = cl.location_id order by l.name`

	rows, err := m.db.Query(c, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query cinema: %w", err)
	}
	defer rows.Close()

	var cinemas []models.Cinema
	for rows.Next() {
		var cinema models.Cinema
		if err := rows.Scan(&cinema.Id, &cinema.Name, &cinema.Location); err != nil {
			return nil, fmt.Errorf("failed to scan cinema: %w", err)
		}
		cinemas = append(cinemas, cinema)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return cinemas, nil
}
