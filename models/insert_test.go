package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsertCotacao(t *testing.T) {
	arg := Cambio{
		Moeda_Inicial: "BRL",
		Moeda_Final:   "USD",
		Cotacao:       05.10,
		Valor_Inicial: 100.00,
	}

	err := InsertCotacao(arg)
	require.NoError(t, err)
}
