package models

type ProductCreateInput struct {
	CategoryID  *int     `json:"categoryID"`
	Sku         *string  `json:"sku"`
	Name        *string  `json:"name"`
	Stock       *int     `json:"stock"`
	Description *string  `json:"description"`
	Price       *float32 `json:"price"`
	Discount    *float32 `json:"discount"`
	Status      *int8    `json:"status"`
}

type ProductUpdateInput struct {
	CategoryID  *int     `json:"categoryID"`
	Sku         *string  `json:"sku"`
	Name        *string  `json:"name"`
	Stock       *int     `json:"stock"`
	Description *string  `json:"description"`
	Price       *float32 `json:"price"`
	Discount    *float32 `json:"discount"`
	Status      *int8    `json:"status"`
}
