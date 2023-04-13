package client

import (
	"BonsaiStore/functions"
	"BonsaiStore/structs"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateClients(c *gin.Context, _ *sql.DB) {
	firstname, lastname, sex := functions.RandomPersonName()
	name := firstname + " " + lastname + " " + sex
	response := structs.Response{Status: "Clients generated " + name}
	c.JSON(http.StatusOK, response)
}

func DeleteClients(c *gin.Context, db *sql.DB) {
	_, err := db.Exec("delete from cliente")
	if err != nil {
		panic(err)
	}
	response := structs.Response{Status: "Clients deleted"}
	c.JSON(http.StatusOK, response)
}
