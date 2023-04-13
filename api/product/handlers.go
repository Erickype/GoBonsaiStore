package product

import (
	"BonsaiStore/functions"
	"BonsaiStore/structs"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateBonsaiProducts(c *gin.Context, db *sql.DB) {
	bonsaiCategories := getCategoriesByParentId(db, 1)
	nProducts := functions.NewRandomValue(500)
	query := "insert into producto (categoria_id, nombre, fecha_fabricacion)  values ($1, $2,$3)"
	for i := 0; i < nProducts; i++ {
		name := functions.RandomFancyName()
		catId := functions.RandomIndexValue(bonsaiCategories)
		date := functions.GenerateDateInRange(1980, 2020)
		_, err := db.Exec(query, catId, name, date)
		if err != nil {
			panic(err)
		}
	}
	response := structs.Response{Status: "ok"}
	c.JSON(http.StatusOK, response)
}

func DeleteBonsaiProducts(c *gin.Context, db *sql.DB) {
	_, err := db.Exec("delete from producto")
	if err != nil {
		panic(err)
	}
	fmt.Printf("BonsaiProducts deleted")
	response := structs.Response{Status: "ok"}
	c.JSON(http.StatusOK, response)
}
