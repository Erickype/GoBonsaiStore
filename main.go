package main

import (
	"BonsaiStore/api/country"
	"BonsaiStore/api/product"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	db, err := sql.Open("postgres", "user=postgres dbname=BonsaiStore password=Erickype host=localhost port=5432 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	router := gin.Default()

	country.Routes(router, db)
	product.Routes(router, db)

	err = router.Run()
	if err != nil {
		panic(err)
	}
}
