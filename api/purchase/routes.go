package purchase

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, db *sql.DB) {
	router.GET("/purchases/generate", func(context *gin.Context) {
		GeneratePurchases(context, db)
	})
	router.GET("/purchases/delete", func(context *gin.Context) {
		DeletePurchases(context, db)
	})
}
