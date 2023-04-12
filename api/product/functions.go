package product

import (
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
		var category int
		err := rows.Scan(&category)
		if err != nil {
			log.Fatal(err)
		}
		categoriesId = append(categoriesId, category)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return categoriesId
}
