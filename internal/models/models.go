package models

import "time"

type QueryParams struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Query  string `json:"query"`
}

type User struct {
	ID         int       `json:"id"`
	UserTypeID int       `json:"userTypeID"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Avatar     string    `json:"avatar"`
	Phone      string    `json:"phone"`
	Status     int       `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type UserType struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
