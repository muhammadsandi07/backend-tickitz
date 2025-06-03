package models

import "time"

type OrderStruct struct {
	Id          int       `json:"id,omitempty" db:"id"`
	IdUser      int       `json:"userid,omitempty" db:"id_user"`
	IdSchedule  int       `json:"idSchedule" db:"id_schedule"`
	IdSeat      []string  `json:"seat" db:"id_seat"`
	IdPayment   int       `json:"idmethod,omitempty" db:"idmethod"`
	Fullname    string    `json:"fullname" db:"fullname"`
	PhoneNumber string    `json:"phonenumber" db:"phonenumber"`
	IsPaid      bool      `json:"isPaid" db:"id"`
	NameCinema  string    `json:"namecinema" db:"namecinema"`
	Date        time.Time `json:"date"`
	TotalPrice  int       `json:"total_price" db:"total"`
	NameMovie   string    `json:"namemovie" db:"namemovie"`
}
