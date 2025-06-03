package models

type MovieQueryParams struct {
	Page     int    `json:"page" form:"page" binding:"numeric"`
	Name     string `json:"name" form:"name"`
	Genre    string `json:"genre" form:"genre"`
	Location string `json:"location" form:"location"`
}
type ProductQueryParams struct {
	Page     int    `json:"page" form:"page" binding:"numeric"`
	Search   string `json:"search" form:"search"`
	Genre    string `json:"options" form:"options"`
	Discount string `json:"discount" from:"discount"`
}
