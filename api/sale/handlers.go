package sale

import (
	"BonsaiStore/functions"
	"BonsaiStore/structs"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateSales(c *gin.Context, _ *sql.DB) {
	nPurchases := functions.NewRandomValue(5000)

	message := fmt.Sprintf("Sales generated: %d", nPurchases)
	response := structs.Response{Status: message}
	c.JSON(http.StatusOK, response)
}

func GenerateSalesDetails(c *gin.Context, _ *sql.DB) {
	nPurchases := functions.NewRandomValue(1)
	message := fmt.Sprintf("Sales details generated: %d", nPurchases)
	response := structs.Response{Status: message}
	c.JSON(http.StatusOK, response)
}

func DeleteSalesDetails(c *gin.Context, db *sql.DB) {
	exec, err := db.Exec("delete from venta_detalle")
	if err != nil {
		panic(err)
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		panic(err)
	}

	message := fmt.Sprintf("Sales details deleted: %d", affected)
	response := structs.Response{Status: message}
	c.JSON(http.StatusOK, response)
}

func DeleteSales(c *gin.Context, db *sql.DB) {
	exec, err := db.Exec("delete from venta")
	if err != nil {
		panic(err)
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		panic(err)
	}

	message := fmt.Sprintf("Sales deleted: %d", affected)
	response := structs.Response{Status: message}
	c.JSON(http.StatusOK, response)
}
