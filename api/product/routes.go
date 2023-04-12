package product

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, db *sql.DB) {
	router.GET("/products/bonsai/generate", func(context *gin.Context) {
		GenerateBonsaiProducts(context, db)
	})
	router.GET("/products/bonsai/delete", func(context *gin.Context) {
		DeleteBonsaiProducts(context, db)
	})
}
