package model

type Response struct {
	Id          int         `json:"id"`
	Quantidade  int         `json:"quantidade"`
	Tipo        Tipo        `json:"tipo"`
	Card        []Card      `json:"card"`
	Estatistica Estatistica `json:"estatistica"`
}
