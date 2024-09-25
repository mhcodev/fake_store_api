package models

import "time"

type User struct {
	ID         int
	UserTypeID int
	Name       string
	Email      string
	Password   string
	Avatar     string
	Phone      string
	Status     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
