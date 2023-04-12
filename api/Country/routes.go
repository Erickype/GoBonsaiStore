package Country

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, db *sql.DB) {
	router.GET("/countries", func(context *gin.Context) {
		GetCountries(context, db)
	})
	router.GET("/countries/generate", func(context *gin.Context) {
		GenerateCountries(context, db)
	})
}
