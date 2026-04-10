package model

type Pokemon struct {
	Id         int32        `json:"id"`
	Nome       string       `json:"nome"`
	TipoCarta  string       `json:"card_type"`
	Tipo       Tipo         `json:"tipo"`
	Estagio    int32        `json:"estagio"`
	Habilidade []Habilidade `json:"habilidade"`
	Ataque     []Ataque     `json:"ataque"`
	Geracao    int32        `json:"geracao"`
	PS         int32        `json:"ps"`
	Recuo      int32        `json:"recuo"`
	Fraqueza   Tipo         `json:"fraqueza"`
}
