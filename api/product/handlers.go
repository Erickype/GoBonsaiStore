package product

import (
	"BonsaiStore/functions"
	"BonsaiStore/generators"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateBonsaiProducts(c *gin.Context, db *sql.DB) {
	bonsaiCategories := getCategoriesByParentId(db, 1)
	nProducts := functions.NewRandomValue(500)
	query := "insert into producto (categoria_id, nombre, fecha_fabricacion)  values ($1, $2,$3)"
	for i := 0; i < nProducts; i++ {
		name := generators.GenerateRandomName()
		catId := functions.RandomIndex(bonsaiCategories)
		date := functions.GenerateDateRange().Format("2006-01-02 15:04:05.999999-07:00")
		_, err := db.Exec(query, catId, name, date)
		if err != nil {
			c.JSON(http.StatusConflict, bonsaiCategories)
		}
	}
	c.JSON(http.StatusOK, bonsaiCategories)
}
