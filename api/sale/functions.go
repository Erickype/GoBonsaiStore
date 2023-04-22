package sale

import (
	"BonsaiStore/api/inventory"
	"BonsaiStore/functions"
	"database/sql"
	"log"
)

func Count(db *sql.DB) int {
	var rows int
	query := "select count(venta_id) from venta"
	row := db.QueryRow(query)
	err := row.Scan(&rows)
	if err != nil {
		panic(err)
	}
	return rows
}

func GetIds(db *sql.DB) []int {
	rows, err := db.Query("SELECT venta_id FROM venta where cancelado = false")
	if err != nil {
		panic(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var ids []int

	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	return ids
}

func generateDetail(db *sql.DB, productId int, saleId int) *Detail {
	detail := &Detail{
		ProductId: productId,
		SaleId:    saleId,
	}

	inventoryProduct := inventory.ProductById(db, productId)

	quantity := functions.NewRandomValue(5) + 1
	if quantity > inventoryProduct.Stock {
		quantity = functions.NewRandomValue(inventoryProduct.Stock) + 1
	}

	detail.ProductQuantity = quantity

	return detail
}
