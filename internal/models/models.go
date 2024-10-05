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
	Status     int8      `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type UserType struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Category struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"imageURL"`
	Status   int8   `json:"status"`
}

type Product struct {
	ID          int       `json:"id"`
	CategoryID  int       `json:"categoryID"`
	Sku         string    `json:"sku"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Stock       int       `json:"stock"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
	Discount    float32   `json:"discount"`
	Status      int8      `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
