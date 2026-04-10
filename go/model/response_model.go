package model

type Response struct {
	Id          int32       `json:"id"`
	Quantidade  int32       `json:"quantidade"`
	Tipo        Tipo        `json:"tipo"`
	Card        []Card      `json:"card"`
	Estatistica Estatistica `json:"estatistica"`
}
