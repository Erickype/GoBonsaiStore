package product

import (
	"BonsaiStore/api/category"
	"BonsaiStore/functions"
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

func GeneratePriceQuantity(db *sql.DB, productId int) (float64, int) {
	var price float64
	var quantity int
	product := GetById(db, productId)
	fatherCategory := category.GetIdFromChild(db, product.CategoryId)
	switch fatherCategory {
	case 1:
		quantity = functions.NewRandomValue(10) + 1
		age := AgeInYears(db, productId)
		basePrice := prices[1]
		price = float64(age) * basePrice
		break
	case 12:
		quantity = functions.NewRandomValue(30) + 1
		price = prices[12]
		break
	case 13:
		quantity = functions.NewRandomValue(10) + 1
		price = prices[13]
	case 14:
		var age int
		basePrice := prices[1]
		quantity = functions.NewRandomValue(10) + 1
		if product.CategoryId == 33 {
			age = AgeInYears(db, productId)
			price = float64(age) * basePrice
		} else {
			price = basePrice
		}
	}
	return price, quantity
}

func GenerateMeasuresPercentageOffset(category, subcategory int) (randomTops []float64) {
	stdMeasures := measures[category][subcategory]
	percentage := functions.NewRandomValue(50) + 1
	var offsets []float64
	for i, measure := range stdMeasures {
		offset := measure * float64(percentage) / 100
		offsets = append(offsets, offset)
		tails := functions.NewRandomValue(2)
		if tails == 0 {
			randomTop := stdMeasures[i] + offsets[i]
			randomTops = append(randomTops, randomTop)
		} else {
			randomTop := stdMeasures[i] - offsets[i]
			randomTops = append(randomTops, randomTop)
		}
	}
	return
}
