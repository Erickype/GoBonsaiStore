package supplier

import (
	"BonsaiStore/api/country"
	"BonsaiStore/functions"
	"BonsaiStore/structs"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateSuppliers(c *gin.Context, db *sql.DB) {
	countriesId := country.GetIds(db)
	nSuppliers := functions.NewRandomValue(100)
	query := "insert into proveedor (pais_id, nombre, contacto, direccion, evaluacion) values ($1,$2,$3,$4,$5)"

	for i := 0; i < nSuppliers; i++ {
		countryId := functions.RandomIndexValue(countriesId)
		name := functions.RandomFancyName()
		contact := functions.RandomPhoneWithId(countryId)
		address := functions.RandomAddress()
		evaluation := functions.NewRandomValue(30) + 70

		_, err := db.Exec(query, countryId, name, contact, address, evaluation)
		if err != nil {
			panic(err)
		}
	}

	response := structs.Response{Status: "Suppliers generated"}
	c.JSON(http.StatusOK, response)
}

func DeleteSuppliers(c *gin.Context, db *sql.DB) {
	_, err := db.Exec("delete from proveedor")
	if err != nil {
		panic(err)
	}
	response := structs.Response{Status: "Suppliers deleted"}
	c.JSON(http.StatusOK, response)
}
