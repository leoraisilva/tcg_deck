package model

type EditCard struct {
	Id    int32  `json:"id_deck"`
	Cards []Card `json:"cards"`
}
