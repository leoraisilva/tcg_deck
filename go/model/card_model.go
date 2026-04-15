package model

type Card struct {
	Id       int32  `json:"id"`
	CardType string `json:"card_tipo"`
	Pokemon  int32  `json:"pokemon"`
	Apoiador int32  `json:"apoiador"`
	Item     int32  `json:"item"`
}
