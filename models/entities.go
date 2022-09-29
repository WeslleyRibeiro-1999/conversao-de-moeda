package models

type Cambio struct {
	Moeda_Inicial string  `json:"moeda_inicial"`
	Moeda_Final   string  `json:"moeda_final"`
	Cotacao       float64 `json:"cotacao"`
	Valor_Inicial float64 `json:"valor_inicial"`
}
