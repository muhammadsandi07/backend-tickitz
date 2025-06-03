package models

import (
	"time"
)

type ScheduleStruct struct {
	Id        int       `json:"id,omitempty"  db:"id"`
	MovieName string    `json:"movieid,omitempty" db:"name"`
	Date      time.Time `json:"date,omitempty" db:"date"`
	Price     int       `json:"price,omitempty" db:"price"`
	Cinema    string    `json:"cinema,omitempty" db:"cinema"`
	Location  string    `json:"location,omitempty" db:"location"`
}
