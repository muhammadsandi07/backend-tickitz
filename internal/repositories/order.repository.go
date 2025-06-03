package repositories

import (
	"backendtickitz/internal/models"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{db: db}
}

func (o *OrderRepository) AddOrder(c context.Context, newOrder *models.OrderStruct) error {
	tx, err := o.db.Begin(c)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback(c)
		}
	}()

	query := `insert into orders (id_user, id_schedule, id_payment,ispaid,fullname,phone_number ) values ($1,$2,$3,$4,$5,$6) returning id`
	newValues := []any{newOrder.IdUser, newOrder.IdSchedule, newOrder.IdPayment, false, newOrder.Fullname, newOrder.PhoneNumber}
	log.Println("values", newValues)
	err = tx.QueryRow(c, query, newValues...).Scan(&newOrder.Id)
	if err != nil {
		return err
	}
	log.Println("error insert order", err)

	queryAsosiation := `insert into orders_seats (orders_id,seats_id) values`
	// building query
	values := []any{newOrder.Id}
	for i, seat := range newOrder.IdSeat {
		if i > 0 {
			queryAsosiation += ","
		}
		queryAsosiation += fmt.Sprintf("($1, $%d)", len(values)+1)
		values = append(values, seat)
	}
	cmd, err := tx.Exec(c, queryAsosiation, values...)
	log.Println("error insert ass", err)

	if err != nil {
		return errors.New("gagal eksekusi query")
	}
	row := cmd.RowsAffected()
	if row == 0 {
		return errors.New("tidak ada data yang terpengaruh ")
	}
	if err := tx.Commit(c); err != nil {
		return err
	}
	return nil
}

func (o *OrderRepository) GetOrderByUser(c context.Context, idUser int) ([]models.OrderStruct, error) {
	query := `select  m.name as namemovie ,o.id_schedule,s.date, o.ispaid, c.name as namecinema ,sum (s.price) as total, array_agg(os.seats_id) from orders o join users u on o.id_user = u.id join schedule s on o.id_schedule = s.id join orders_seats os on os.orders_id  = o.id join cinema c on s.id_cinema  = c.id join movie m  on s.id_movie = m.id where u.id = $1 group by m.id, u.id, s.id,o.id, c."name" `
	rows, err := o.db.Query(c, query, idUser)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var result []models.OrderStruct
	for rows.Next() {
		var order models.OrderStruct
		err := rows.Scan(&order.NameMovie, &order.IdSchedule, &order.Date, &order.IsPaid, &order.NameCinema, &order.TotalPrice, &order.IdSeat)
		if err != nil {
			return nil, err
		}
		result = append(result, order)
	}
	log.Println("DEBUG", result)
	return result, nil
}
