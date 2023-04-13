package client

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, db *sql.DB) {
	router.GET("/clients/generate", func(context *gin.Context) {
		GenerateClients(context, db)
	})
	router.GET("/clients/delete", func(context *gin.Context) {
		DeleteClients(context, db)
	})
}
