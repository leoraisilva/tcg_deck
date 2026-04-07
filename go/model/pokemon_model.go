package model

type Pokemon struct {
	Id         int          `json:"id"`
	Nome       string       `json:"nome"`
	TipoCarta  string       `json:"card_type"`
	Tipo       Tipo         `json:"tipo"`
	Estagio    int          `json:"estagio"`
	Habilidade []Habilidade `json:"habilidade"`
	Ataque     []Ataque     `json:"ataque"`
	Geracao    int          `json:"geracao"`
	PS         int          `json:"ps"`
	Recuo      int          `json:"recuo"`
	Fraqueza   Tipo         `json:"fraqueza"`
}
