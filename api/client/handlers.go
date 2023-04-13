package client

import (
	"BonsaiStore/api/country"
	"BonsaiStore/functions"
	"BonsaiStore/structs"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateClients(c *gin.Context, db *sql.DB) {
	nClients := functions.NewRandomValue(2000)

	query := "insert into cliente (pais_id, nombre, apellido, fecha_nacimiento, correo_electronico) values($1,$2,$3,$4,$5)"

	for i := 0; i < nClients; i++ {
		firstname, lastname, _ := functions.RandomPersonName()
		countryId := functions.RandomIndexValue(country.GetIds(db))
		birthdate := functions.GenerateDateInRange(1960, 2000)
		email := functions.RandomEmailAddress(firstname, lastname)

		_, err := db.Exec(query, countryId, firstname, lastname, birthdate, email)
		if err != nil {
			panic(err)
		}
	}

	message := fmt.Sprintf("Clients generated: %d", nClients)
	response := structs.Response{Status: message}
	c.JSON(http.StatusOK, response)
}

func DeleteClients(c *gin.Context, db *sql.DB) {
	exec, err := db.Exec("delete from cliente")
	if err != nil {
		panic(err)
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		panic(err)
	}

	message := fmt.Sprintf("Clients deleted: %d", affected)
	response := structs.Response{Status: message}
	c.JSON(http.StatusOK, response)
}
