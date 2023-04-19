package purchase

import (
	"BonsaiStore/api/product"
	"database/sql"
	"log"
)

func Count(db *sql.DB) int {
	var rows int
	query := "select count(id) from compra"
	row := db.QueryRow(query)
	err := row.Scan(&rows)
	if err != nil {
		panic(err)
	}
	return rows
}

func GetIds(db *sql.DB) []int {
	rows, err := db.Query("SELECT id FROM compra where compra.cancelado = false")
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

func purchaseById(db *sql.DB, id int) *Purchase {
	purchase := &Purchase{}
	query := "select * from compra where id = $1"
	rows := db.QueryRow(query, id)
	err := rows.Scan(&purchase.PurchaseId, &purchase.SupplierId, &purchase.Date, &purchase.Cancel)
	if err != nil {
		panic(err)
	}
	return purchase
}

func generatePurchaseDetail(db *sql.DB, purchase *Purchase, productId int) *Detail {
	detail := &Detail{
		ProductId:  productId,
		PurchaseId: purchase.PurchaseId,
		SupplierId: purchase.SupplierId,
	}

	price, quantity := product.GeneratePriceQuantity(db, productId)
	detail.ProductPrice = price

	detail.ProductQuantity = quantity

	return detail
}
