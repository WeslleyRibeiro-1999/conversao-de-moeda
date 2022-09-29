package main

import (
	"github.com/WeslleyRibeiro-1999/conversao-de-moeda/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/exchange/:amount/:from/:to/:rate", api.GetCotacao)
	router.Run(":8080")
}
