package models

import (
	"mime/multipart"
	"time"
)

type MovieStruct struct {
	Id           int       `json:"id,omitempty" db:"id"`
	Name         string    `json:"name,omitempty"  db:"name"`
	Duration     int       `json:"duration,omitempty"  db:"duration"`
	Synopsis     string    `json:"synopsis,omitempty"  db:"synopsis"`
	Img_movie    string    `json:"img,omitempty"  db:"img_movie"`
	Backdrop     string    `json:"backdrop,omitempty"  db:"backdrop"`
	Release_Date time.Time `json:"release_date"    db:"release_date"`
	Genre        []string  `json:"genre,omitempty"  db:"genre"`
}

type MovieFrom struct {
	Name         string                `form:"title,omitempty"`
	Duration     int                   `form:"duration,omitempty"`
	Synopsis     string                `form:"synopsis,omitempty"`
	Img_movie    *multipart.FileHeader `form:"image,omitempty"`
	Backdrop     *multipart.FileHeader `form:"backdrop,omitempty"`
	Release_Date time.Time             `form:"releaseDate"`
	Genre        []int                 `form:"genre,omitempty"`
	IdCinema     int                   `form:"idcinema"`
	Price        int                   `form:"price"`
	Director     int                   `form:"director"`
	Cast         int                   `form:"cast"`
	ShowDate     time.Time             `form:"showDate"`
}

type ProductStruct struct {
	Search       string                `form:"name,omitempty"`
	Category     int                   `form:"category,omitempty"`
	Options      string                `form:"options,omitempty"`
	Img_movie    *multipart.FileHeader `form:"img,omitempty"`
	Backdrop     *multipart.FileHeader `form:"backdrop,omitempty"`
	Release_Date time.Time             `form:"release_date"`
	Genre        []int                 `form:"genre,omitempty"`
	IdCinema     int                   `form:"idcinema"`
	Price        int                   `form:"price"`
}

type PaginatedMovieResult struct {
	Data       []MovieStruct `json:"data"`
	Page       int           `json:"page"`
	PageSize   int           `json:"pageSize"`
	TotalItems int           `json:"totalItems"`
	TotalPages int           `json:"totalPages"`
}

type Genre struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type Cinema struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}
