package sale

import (
	"BonsaiStore/api/client"
	"BonsaiStore/api/inventory"
	"BonsaiStore/functions"
	"BonsaiStore/structs"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateSales(c *gin.Context, db *sql.DB) {
	query := "insert into venta (cliente_id, fecha, cancelado) values($1,$2,$3)"
	nPurchases := functions.NewRandomValue(4000)
	clientIds := client.GetIds(db)

	for i := 0; i < nPurchases; i++ {
		clientId := functions.RandomIndexValue(clientIds)
		saleDate := functions.GenerateDateInRange(2001, 2023)

		_, err := db.Exec(query, clientId, saleDate, false)
		if err != nil {
			panic(err)
		}
	}

	message := fmt.Sprintf("Sales generated: %d", nPurchases)
	response := structs.Response{Status: message}
	c.JSON(http.StatusOK, response)
}

func GenerateSalesDetails(c *gin.Context, db *sql.DB) {
	query := "insert into venta_detalle (producto_id, venta_id, cantidad_producto) values($1,$2,$3)"
	nSales := functions.NewRandomValue(Count(db))
	productIds := inventory.GetProductsIds(db)
	salesIds := GetIds(db)

	nDetails := 0

	for i := 0; i < nSales; i++ {
		saleId := functions.RandomIndexValue(salesIds)
		nProducts := functions.NewRandomValue(5) + 1
		nDetails += nProducts
		for j := 0; j < nProducts; j++ {
			productId := functions.RandomIndexValue(productIds)
			detail := generateDetail(db, productId, saleId)

			_, err := db.Exec(query, detail.ProductId, detail.SaleId, detail.ProductQuantity)
			if err != nil {
				panic(err)
			}
		}
	}

	message := fmt.Sprintf("Sales details generated: %d", nDetails)
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
