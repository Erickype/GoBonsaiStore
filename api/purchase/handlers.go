package purchase

import (
	"BonsaiStore/api/supplier"
	"BonsaiStore/functions"
	"BonsaiStore/structs"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GeneratePurchases(c *gin.Context, db *sql.DB) {
	nPurchases := functions.NewRandomValue(5000)

	query := "insert into compra (proveedor_id, fecha, cancelado) values($1,$2,$3)"

	for i := 0; i < nPurchases; i++ {
		supplierId := functions.RandomIndexValue(supplier.GetIds(db))
		date := functions.GenerateDateInRange(2000, 2022)

		_, err := db.Exec(query, supplierId, date, false)
		if err != nil {
			panic(err)
		}
	}

	message := fmt.Sprintf("Purchases generated: %d", nPurchases)
	response := structs.Response{Status: message}
	c.JSON(http.StatusOK, response)
}

func DeletePurchases(c *gin.Context, db *sql.DB) {
	exec, err := db.Exec("delete from compra")
	if err != nil {
		panic(err)
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		panic(err)
	}

	message := fmt.Sprintf("Purchases deleted: %d", affected)
	response := structs.Response{Status: message}
	c.JSON(http.StatusOK, response)
}
