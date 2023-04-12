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

	if countries != nil {
		c.JSON(http.StatusOK, countries)
	}

	c.Status(http.StatusNoContent)
}

func GenerateCountries(c *gin.Context, db *sql.DB) {
	for _, country := range countries {
		_, err := db.Exec("INSERT INTO pais (nombre) VALUES ($1)", country)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Inserted row for %s\n", country)
	}
	response := Response{Status: "ok"}
	c.JSON(http.StatusOK, response)
}

func DeleteCountries(c *gin.Context, db *sql.DB) {
	_, err := db.Exec("delete from pais")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Countries deleted")
	response := Response{Status: "ok"}
	c.JSON(http.StatusOK, response)
}
