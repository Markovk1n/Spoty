package models

type Search struct {
	Search string `form:"search" db:"search"`
}
