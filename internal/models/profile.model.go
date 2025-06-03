package models

type ProfileStruct struct {
	UserId       int     `json:"id,omitempty" form:"id,omitempty" db:"user_id"`
	Firstname    *string `json:"firstname" form:"firstname" db:"firstname"`
	Email        string  `json:"email" form:"email" db:"email"`
	Lastname     *string `json:"lastname" form:"lastname" db:"lastname"`
	PhoneNumber  *string `json:"phonenumber,omitempty" form:"phonenumber" db:"phone_number"`
	Point        *int    `json:"point," form:"point,omitempty" db:"point"`
	IdMember     *int    `json:"idmember," form:"idmember,omitempty" db:"id_member"`
	ProfileImage *string `json:"profileimage," form:"profileimage,omitempty" db:"profile_image"`
}
