package models

import (
	"fmt"

	"github.com/WeslleyRibeiro-1999/conversao-de-moeda/database"
)

func InsertCotacao(cambio Cambio) error {
	conn, err := database.OpenConnection()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	_, err = conn.Exec(fmt.Sprintf("INSERT INTO cambio (Moeda_Inicial, Moeda_Final, Cotacao, Valor_Inicial, Data_Cotacao) VALUES ('%s', '%s', %f, %f, now())",
		cambio.Moeda_Inicial, cambio.Moeda_Final, cambio.Cotacao, cambio.Valor_Inicial))

	if err != nil {
		panic(err)
	}

	return err
}
