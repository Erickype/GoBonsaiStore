package supplier

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, db *sql.DB) {
	router.GET("/suppliers/generate", func(context *gin.Context) {
		GenerateSuppliers(context, db)
	})
	router.GET("/suppliers/delete", func(context *gin.Context) {
		DeleteSuppliers(context, db)
	})
}
