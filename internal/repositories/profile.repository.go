package repositories

import (
	"backendtickitz/internal/models"
	"backendtickitz/pkg"
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProfileRepository struct {
	db *pgxpool.Pool
}

func NewProfileRepostory(db *pgxpool.Pool) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (p *ProfileRepository) GetProfileById(c context.Context, idUser int) (models.ProfileStruct, error) {
	query := `select u.id,p.firstname,p.lastname,p.phone_number, p.point,p.id_member,p.profile_image, u.email from profile p join users u on p.user_id = u.id  where u.id = $1`
	log.Println("query", idUser)
	var result models.ProfileStruct
	if err := pkg.DB.QueryRow(c, query, idUser).Scan(&result.UserId, &result.Firstname, &result.Lastname, &result.PhoneNumber, &result.Point, &result.IdMember, &result.ProfileImage, &result.Email); err != nil {
		return models.ProfileStruct{}, err
	}
	return result, nil

}
func (p *ProfileRepository) UpdateProfile(c context.Context, newProfile *models.ProfileStruct, idUser int) (models.ProfileStruct, error) {
	value := []any{newProfile.Firstname, newProfile.Lastname, newProfile.PhoneNumber, newProfile.Point, newProfile.IdMember, newProfile.ProfileImage, idUser}
	query := `update profile set firstname =$1, lastname = $2,phone_number = $3, point = $4, id_member=$5, profile_image=$6 where user_id= $7 returning firstname, lastname, phone_number, point, id_member, profile_image`
	var result models.ProfileStruct

	if err := pkg.DB.QueryRow(c, query, value...).Scan(&result.Firstname, &result.Lastname, &result.PhoneNumber, &result.Point, &result.IdMember, &result.ProfileImage); err != nil {
		return models.ProfileStruct{}, err
	}
	return result, nil

}
