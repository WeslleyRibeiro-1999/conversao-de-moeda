package api

import (
	"net/http"
	"strconv"

	"github.com/WeslleyRibeiro-1999/conversao-de-moeda/models"
	"github.com/gin-gonic/gin"
)

var moedas = map[string]string{
	"BRL":  "R$",
	"USD":  "$",
	"EURO": "€",
	"BTC":  "BTC",
}
var simbolo string

func GetCotacao(c *gin.Context) {
	amount := c.Param("amount")
	from := c.Param("from")
	to := c.Param("to")
	rate := c.Param("rate")

	cotacao, err := strconv.ParseFloat(rate, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"api started with error:": err.Error()})
	}

	valor, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"api started with error:": err.Error()})
	}

	err = models.InsertCotacao(models.Cambio{
		Moeda_Inicial: from,
		Moeda_Final:   to,
		Cotacao:       cotacao,
		Valor_Inicial: valor,
	})

	for i, valor := range moedas {
		if to == i {
			simbolo = valor
		}
	}

	c.JSON(http.StatusOK, gin.H{"valorConvertido": valor * cotacao, "simboloMoeda": simbolo})

}
