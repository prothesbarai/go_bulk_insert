package models

type ProductModel struct{
	Name string `json:"name"`
	Price float64 `json:"price"`
}


type BulkProductRequest struct{
	Products []ProductModel `json:"products"`
}