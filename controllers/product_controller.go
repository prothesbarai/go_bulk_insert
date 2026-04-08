package controllers

import (
	"database/sql"
	"go_bulk_insert/database"
	"go_bulk_insert/logger"
	"go_bulk_insert/models"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

var batchSize = 500

func BulkInsertProducts(ginCtx *gin.Context) {
	var reqBulkPM models.BulkProductRequest

	/// >>>  JSON Binding
	if err := ginCtx.ShouldBindJSON(&reqBulkPM); err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid JSON Format"})
		logger.AppLogger.Error.Println("Invalid JSON Format")
		return
	}

	/// >>> Safety Check
	if len(reqBulkPM.Products) == 0 {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "Empty Product List"})
		logger.AppLogger.Error.Println("Empty Product List")
		return
	}

	/// >>> Extra Safety Check
	if len(reqBulkPM.Products) > 1000 {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "Max 1000 products per request allowed"})
		logger.AppLogger.Error.Println("Max 1000 products per request allowed")
		return
	}

	/// >>> Database Transection
	tx, err := database.DB.Begin()
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
		logger.AppLogger.Error.Println("Failed to start transaction")
		return
	}

	totalInserted := 0

	// >>> Batch processing
	for i := 0; i < len(reqBulkPM.Products); i = i + batchSize {
		end := i + batchSize // >>> 1st Batch 0 + 500 = 500 means 0-499 index
		if end > len(reqBulkPM.Products) {
			end = len(reqBulkPM.Products)
		}
		/*
			Product = 2000
			i = 0,     batchSize = 500 →   end = 0 +  500 = 500      → batch index [0:500] = 0–499
			i = 500,   batchSize = 500 →   end = 500 + 500 = 1000    → batch index [500:1000] = 500–999
			i = 1000,  batchSize = 500 →   end = 1000 + 500 = 1500   → batch index [1000:1500] = 1000–1499
			i = 1500,  batchSize = 500 →   end = 1500 + 500 = 2000   → batch index [1500:2000] = 1500–1999

		*/
		batch := reqBulkPM.Products[i:end]

		if err := insertBatch(tx, batch); err != nil {
			tx.Rollback() // When 1 Batch Error Then All Cancel Purpose
			ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "Batch Insetion Failed"})
			logger.AppLogger.Error.Println("Batch Insetion Failed : ", err)
			return
		}

		totalInserted = totalInserted + len(batch)
	}

	// >>> Commit transaction
	if err := tx.Commit(); err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "Commit failed"})
		logger.AppLogger.Error.Println("Commit failed")
		return
	}

	/// >>> If All Success
	ginCtx.JSON(http.StatusOK, gin.H{"message": "Bulk insert success", "inserted": totalInserted})
	logger.AppLogger.Info.Println("Bulk insert success. Total-Inseted : ", totalInserted)

}

func insertBatch(tx *sql.Tx, products []models.ProductModel) error {
	query := "INSERT INTO products (name, store_id, store_code, category_id, subcategory_id, sub_subcategory_id, photos, thumbnail, featured_img, video_link, tags, description, price, purchase_price, discount, discount_type, discounted_price,sku, unit, weight, variant_product, attributes, choice_options, colors, variations, published, trashed, stock_in, featured, created_by, created_at, updated_at) VALUES "

	palceholder := make([]string, 0, len(products))
	values := make([]interface{}, 0, len(products)*2)

	for _, p := range products {
		// >>> Basic validation (safe)
		if p.Name == "" {
			p.Name = "Unnamed"
		}
		palceholder = append(palceholder, "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
		values = append(values, p.Name)
		values = append(values, p.StoreId)
		values = append(values, p.StoreCode)
		values = append(values, p.CategoryId)
		values = append(values, p.SubCategoryId)
		values = append(values, p.SubSubCategoryId)
		values = append(values, p.Photos)
		values = append(values, p.Thumbnail)
		values = append(values, p.FeaturedImg)
		values = append(values, p.VideoLink)
		values = append(values, p.Tags)
		values = append(values, p.Description)
		values = append(values, p.Price)
		values = append(values, p.PurchasePrice)
		values = append(values, p.Discount)
		values = append(values, string(p.DiscountType))
		values = append(values, p.DiscountedPrice)
		values = append(values, p.Sku)
		values = append(values, p.Unit)
		values = append(values, p.Weight)
		values = append(values, p.VariantProduct)
		values = append(values, p.Attributes)
		values = append(values, p.ChoiceOptions)
		values = append(values, p.Colors)
		values = append(values, p.Variations)
		values = append(values, p.Published)
		values = append(values, p.Trashed)
		values = append(values, p.StockIn)
		values = append(values, p.Featured)
		values = append(values, p.CreatedBy)
		values = append(values, p.CreatedAt)
		values = append(values, p.UpdatedAt)
	}

	query += strings.Join(palceholder, ",")
	_, err := tx.Exec(query, values...)
	return err
}
