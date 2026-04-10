package model

type Apoiador struct {
	Id       int32  `json:"id"`
	Nome     string `json:"nome"`
	CardType string `json:"card_type"`
	Efeito   string `json:"efeito"`
}
