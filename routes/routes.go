package routes

import (
	"go_bulk_insert/controllers"
	"github.com/gin-gonic/gin"
)

func SetRoutes(route *gin.Engine){
	route.POST("products/bulk",controllers.BulkInsertProducts)
}