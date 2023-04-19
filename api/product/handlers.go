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

func GenerateSoilProducts(c *gin.Context, db *sql.DB) {
	soilCategories := getCategoriesByParentId(db, 12)
	nProducts := functions.NewRandomValue(500)
	query := "insert into producto (categoria_id, nombre, fecha_fabricacion)  values ($1, $2,$3)"
	for i := 0; i < nProducts; i++ {
		name := functions.RandomFancyName()
		catId := functions.RandomIndexValue(soilCategories)
		date := functions.GenerateDateInRange(2020, 2022)
		_, err := db.Exec(query, catId, name, date)
		if err != nil {
			panic(err)
		}
	}
	response := structs.Response{Status: "Soil products generated"}
	c.JSON(http.StatusOK, response)
}

func GenerateToolsProducts(c *gin.Context, db *sql.DB) {
	soilCategories := getCategoriesByParentId(db, 13)
	nProducts := functions.NewRandomValue(500)
	query := "insert into producto (categoria_id, nombre, fecha_fabricacion)  values ($1, $2,$3)"
	for i := 0; i < nProducts; i++ {
		name := functions.RandomFancyName()
		catId := functions.RandomIndexValue(soilCategories)
		date := functions.GenerateDateInRange(2020, 2022)
		_, err := db.Exec(query, catId, name, date)
		if err != nil {
			panic(err)
		}
	}
	response := structs.Response{Status: "Tools products generated"}
	c.JSON(http.StatusOK, response)
}

func GeneratePotsProducts(c *gin.Context, db *sql.DB) {
	soilCategories := getCategoriesByParentId(db, 14)
	nProducts := functions.NewRandomValue(500)
	query := "insert into producto (categoria_id, nombre, fecha_fabricacion)  values ($1, $2,$3)"
	for i := 0; i < nProducts; i++ {
		name := functions.RandomFancyName()
		var date string
		catId := functions.RandomIndexValue(soilCategories)
		if catId == 33 {
			date = functions.GenerateDateInRange(1980, 2000)
		} else {
			date = functions.GenerateDateInRange(2020, 2022)
		}
		_, err := db.Exec(query, catId, name, date)
		if err != nil {
			panic(err)
		}
	}
	response := structs.Response{Status: "Pots products generated"}
	c.JSON(http.StatusOK, response)
}

func DeleteBonsaiProducts(c *gin.Context, db *sql.DB) {
	_, err := db.Exec("delete from producto as P\nwhere P.categoria_id in (select Sub.categoria_id\n                     from categoría as Sub\n                     where Sub.categoria_padre = (select Cat.categoria_id\n                                                  from categoría as Cat\n                                                  where Cat.nombre = 'Bonsai'));\n")
	if err != nil {
		panic(err)
	}
	fmt.Printf("BonsaiProducts deleted")
	response := structs.Response{Status: "ok"}
	c.JSON(http.StatusOK, response)
}
