package repositories

import (
	"backendtickitz/internal/models"
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ScheduleRepository struct {
	db *pgxpool.Pool
}

func NewScheduleRepository(db *pgxpool.Pool) *ScheduleRepository {
	return &ScheduleRepository{db: db}
}

func (s *ScheduleRepository) GetSchedule(c context.Context, idMovie int, params string) ([]models.ScheduleStruct, error) {
	query := `select s.id, m.name as name, c.name as cinema, l.name as location, s.price, s.date from schedule s join movie m on s.id_movie  = m.id join cinema_location cl on s.id_cinema = cl.id join cinema c on cl.id = c.id join location l on cl.location_id  = l.id  where m.id= $1 `
	newValues := []any{idMovie}
	if params != "" {
		query += ` and l.name = $2`
		newValues = append(newValues, params)
	}
	query += " group by m.id, s.id, c.id, l.id"

	rows, err := s.db.Query(c, query, newValues...)
	if err != nil {
		log.Println("error schedule", err)
		return nil, err
	}
	log.Println("anak ayam")
	defer rows.Close()
	var result []models.ScheduleStruct
	for rows.Next() {
		var schedule models.ScheduleStruct
		log.Println("ini ini")
		if err := rows.Scan(&schedule.Id, &schedule.MovieName, &schedule.Cinema, &schedule.Location, &schedule.Price, &schedule.Date); err != nil {
			log.Println("ini error price", err)
			return nil, err
		}
		log.Println("schedule ini ", schedule)
		result = append(result, schedule)
	}
	log.Println("ini result schedule", result)
	return result, nil
}
