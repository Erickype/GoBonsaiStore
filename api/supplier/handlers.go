package supplier

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateSuppliers(c *gin.Context, db *sql.DB) {
	response := Response{Status: "Suppliers generated"}
	c.JSON(http.StatusOK, response)
}

func DeleteSuppliers(c *gin.Context, db *sql.DB) {
	_, err := db.Exec("delete from producto")
	if err != nil {
		panic(err)
	}
	response := Response{Status: "Suppliers deleted"}
	c.JSON(http.StatusOK, response)
}
