package sale

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, db *sql.DB) {
	router.GET("/sales/generate", func(context *gin.Context) {
		GenerateSales(context, db)
	})
	router.GET("/sales/delete", func(context *gin.Context) {
		DeleteSales(context, db)
	})
	router.GET("/sales/details/generate", func(context *gin.Context) {
		GenerateSalesDetails(context, db)
	})
	router.GET("/sales/details/delete", func(context *gin.Context) {
		DeleteSalesDetails(context, db)
	})
}
