package repositories

import (
	"backendtickitz/internal/models"
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{db: db}
}

func (a *AuthRepository) Register(c context.Context, newUser *models.AuthStruct, hashed string) (*models.ResultCommand, error) {

	tx, err := a.db.Begin(c)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(c)
		}
	}()

	queryIsEmailExist := "select email from users where email = $1"
	cmd, err := tx.Exec(c, queryIsEmailExist, newUser.Email)
	if err != nil {
		log.Println("data is exist ", err)

		return nil, err
	}
	isExist := cmd.RowsAffected()
	if isExist != 0 {
		return nil, errors.New("email is exist")
	}
	var resultId models.AuthStruct
	query := "insert into users (email,password) values ($1,$2) returning id"
	errinsert := tx.QueryRow(c, query, newUser.Email, hashed).Scan(&resultId.Id)
	if errinsert != nil {
		return nil, errinsert
	}
	log.Println("id", resultId)
	queryProfile := "insert into profile (user_id) values ($1)"
	_, errProfile := tx.Exec(c, queryProfile, resultId.Id)
	if errProfile != nil {
		return nil, errProfile
	}
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(c); err != nil {
		return nil, err
	}
	return &models.ResultCommand{
		ResultFirst:  pgconn.CommandTag{},
		ResultSecond: pgconn.CommandTag{},
	}, nil
}

func (a *AuthRepository) Login(c context.Context, dataUser *models.AuthStruct) (models.AuthStruct, error) {
	query := "select id,password,role from users where email = $1"
	var dataDb models.AuthStruct
	if err := a.db.QueryRow(c, query, dataUser.Email).Scan(&dataDb.Id, &dataDb.Password, &dataDb.Role); err != nil {
		return models.AuthStruct{}, err
	}
	log.Println("password dari db", dataDb.Password)
	return dataDb, nil
}
