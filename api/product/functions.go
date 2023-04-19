package product

import (
	"BonsaiStore/api/category"
	"database/sql"
	"log"
)

func getCategoriesByParentId(db *sql.DB, parentId int) []int {
	rows, err := db.Query("SELECT categoria_id FROM categoría where categoria_padre = $1", parentId)
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var categoriesId []int

	for rows.Next() {
		var categoryId int
		err := rows.Scan(&categoryId)
		if err != nil {
			log.Fatal(err)
		}
		categoriesId = append(categoriesId, categoryId)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return categoriesId
}

func GetIds(db *sql.DB) []int {
	rows, err := db.Query("SELECT producto_id FROM producto")
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

func GetById(db *sql.DB, productId int) *Product {
	product := &Product{}
	query := "select * from producto where producto_id = $1"

	rows := db.QueryRow(query, productId)

	err := rows.Scan(&product.Id, &product.CategoryId, &product.Name, &product.ManufacturingDate)
	if err != nil {
		panic(err)
	}

	return product
}

func AgeInYears(db *sql.DB, productId int) int {
	var age int
	query := "SELECT date_part('year',age(fecha_fabricacion)) AS age FROM producto WHERE producto_id = $1"

	rows := db.QueryRow(query, productId)

	err := rows.Scan(&age)
	if err != nil {
		panic(err)
	}

	return age
}

func GeneratePrice(db *sql.DB, productId int) float64 {
	var price float64
	product := GetById(db, productId)
	fatherCategory := category.GetIdFromChild(db, product.CategoryId)
	switch fatherCategory {
	case 1:
		age := AgeInYears(db, productId)
		basePrice := prices[1]
		price = float64(age) * basePrice
		break
	case 12:
		price = prices[12]
		break
	case 13:
		price = prices[13]
	case 14:
		age := AgeInYears(db, productId)
		basePrice := prices[1]
		price = float64(age) * basePrice
	}
	return price
}
