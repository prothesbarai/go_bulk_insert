package models

type ProductModel struct {
	Name             string    `json:"name"`
	StoreId          int       `json:"store_id"`
	StoreCode        string    `json:"store_code"`
	CategoryId       string    `json:"category_id"`
	SubCategoryId    string    `json:"subcategory_id"`
	SubSubCategoryId string    `json:"sub_subcategory_id"`
	Photos           string    `json:"photos"`
	Thumbnail        string    `json:"thumbnail"`
	FeaturedImg      string    `json:"featured_img"`
	VideoLink        string    `json:"video_link"`
	Tags             string    `json:"tags"`
	Description      string    `json:"description"`
	Price            float64   `json:"price"`
	PurchasePrice    float64   `json:"purchase_price"`
	Discount         float64   `json:"discount"`
	DiscountType     DiscountType    `json:"discount_type"`  // >>> This is Enum Type So Need Helper Model custom_type_model.go
	DiscountedPrice  float64   `json:"discounted_price"`
	Sku              string    `json:"sku"`
	Unit             string    `json:"unit"`
	Weight           float64   `json:"weight"`
	VariantProduct   int       `json:"variant_product"`
	Attributes       string    `json:"attributes"`
	ChoiceOptions    string    `json:"choice_options"`
	Colors           string    `json:"colors"`
	Variations       string    `json:"variations"`
	Published        int       `json:"published"`
	Trashed          int       `json:"trashed"`
	StockIn          int       `json:"stock_in"`
	Featured         int       `json:"featured"`
	CreatedBy        int       `json:"created_by"`
	CreatedAt        string    `json:"created_at"`
	UpdatedAt        string    `json:"updated_at"`
}

type BulkProductRequest struct {
	Products []ProductModel `json:"products"`
}