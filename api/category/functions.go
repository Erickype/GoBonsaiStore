package category

import (
	"database/sql"
	"log"
)

func GetChildIds(db *sql.DB) []int {
	rows, err := db.Query("SELECT categoria_id FROM categoría where categoria_padre is not null ")
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

func GetIdFromChild(db *sql.DB, childId int) int {
	var id int
	query := "select categoria_padre from categoría where categoria_id = $1"

	row := db.QueryRow(query, childId)

	err := row.Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}
