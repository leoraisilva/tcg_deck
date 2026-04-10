package model

type Card struct {
	Id       int32    `json:"id"`
	CardType string   `json:"card_tipo"`
	Pokemon  Pokemon  `json:"pokemon"`
	Apoiador Apoiador `json:"apoiador"`
	Item     Item     `json:"item"`
}
