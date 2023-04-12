package Country

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetCountries(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT * FROM pais")
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var countries []Country

	for rows.Next() {
		var country Country
		err := rows.Scan(&country.Id, &country.Name)
		if err != nil {
			log.Fatal(err)
		}
		countries = append(countries, country)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, countries)
}

func GenerateCountries(_ *gin.Context, db *sql.DB) {
	for _, country := range countries {
		_, err := db.Exec("INSERT INTO pais (nombre) VALUES ($1)", country)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Inserted row for %s\n", country)
	}
}
