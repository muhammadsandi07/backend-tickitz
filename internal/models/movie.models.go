package models

import "time"

type MovieStruct struct {
	Id           int       `json:"id,omitempty" form:"id" db:"id"`
	Name         string    `json:"name" form:"name" db:"name"`
	Duration     int       `json:"duration" form:"duration" db:"duration"`
	Synopsis     string    `json:"synopsis" form:"synopsis" db:"synopsis"`
	Img_movie    string    `json:"img_movie" form:img_movie db:"img_movie"`
	Backdrop     string    `json:"backdrop" form:backdrop db:"backdrop"`
	Release_Date time.Time `json:"release_date" form:release_date db:"release_date"`
	Genre        []any     `json:"genre" form:genre db:"genre"`
}
