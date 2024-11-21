package models

import "time"

type QueryParams struct {
	Limit     int                    `json:"limit"`
	Offset    int                    `json:"offset"`
	Query     string                 `json:"query"`
	MapParams map[string]interface{} `json:"params"`
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
	Images      []string  `json:"images"`
	Category    Category  `json:"category"`
	Status      int8      `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Order struct {
	ID              int       `json:"id"`
	UserID          int       `json:"userID"`
	UserEmail       string    `json:"userEmail"`
	UserName        string    `json:"userName"`
	Quantity        float32   `json:"quantity"`
	Subtotal        float32   `json:"subtotal"`
	DiscountTotal   float32   `json:"discountTotal"`
	Total           float32   `json:"total"`
	ShippingAddress string    `json:"ShippingAddress"`
	Status          int8      `json:"status"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type OrderDetail struct {
	ID               int       `json:"id"`
	OrderID          int       `json:"orderID"`
	ProductID        string    `json:"productID"`
	ProductSku       string    `json:"productSKU"`
	CategoryName     string    `json:"categoryName"`
	ProductName      string    `json:"productName"`
	ProductPriceUnit string    `json:"productPriceUnit"`
	ProductQuantity  string    `json:"productQuantity"`
	ProductSubtotal  string    `json:"productSubtotal"`
	Status           int8      `json:"status"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

type File struct {
	ID           int       `json:"id"`
	OriginalName string    `json:"originalName"`
	FileName     string    `json:"filename"`
	Type         string    `json:"type"`
	Url          string    `json:"url"`
	Status       int8      `json:"status"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
