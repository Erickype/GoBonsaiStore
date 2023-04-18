package purchase

import (
	"BonsaiStore/api/product"
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

func GeneratePurchasesDetails(c *gin.Context, db *sql.DB) {
	nPurchases := functions.NewRandomValue(Count(db))
	purchaseIds := GetIds(db)
	productIds := product.GetIds(db)

	query := "insert into compra_detalle (producto_id, compra_id, proveedor_id, cantidad_producto, precio_producto) values($1,$2,$3,$4,$5)"

	for i := 0; i < nPurchases; i++ {
		nProducts := functions.NewRandomValue(5) + 1

		purchaseId := functions.RandomIndexValue(purchaseIds)
		purchase := purchaseById(db, purchaseId)

		for j := 0; j < nProducts; j++ {
			productId := functions.RandomIndexValue(productIds)
			detail := generatePurchaseDetail(db, purchase, productId)
			_, err := db.Exec(query, detail.ProductId, detail.PurchaseId, detail.SupplierId, detail.ProductQuantity, detail.ProductPrice)
			if err != nil {
				panic(err)
			}
		}
	}

	message := fmt.Sprintf("Purchases details generated: %d", nPurchases)
	response := structs.Response{Status: message}
	c.JSON(http.StatusOK, response)
}

func DeletePurchasesDetails(c *gin.Context, db *sql.DB) {
	exec, err := db.Exec("delete from compra_detalle")
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

func DeletePurchases(c *gin.Context, db *sql.DB) {
	exec, err := db.Exec("delete from compra")
	if err != nil {
		panic(err)
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		panic(err)
	}

	message := fmt.Sprintf("Purchases details deleted: %d", affected)
	response := structs.Response{Status: message}
	c.JSON(http.StatusOK, response)
}
