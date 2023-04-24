package inventory

import (
	"database/sql"
	"log"
)

func GetProductsIds(db *sql.DB) []int {
	rows, err := db.Query("SELECT producto_id FROM inventario where stock_producto > 0")
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

func ProductById(db *sql.DB, id int) *Product {
	product := &Product{}
	query := "select * from inventario where producto_id = $1"
	rows := db.QueryRow(query, id)
	err := rows.Scan(&product.Id, &product.Price, &product.Stock)
	if err != nil {
		panic(err)
	}
	return product
}
