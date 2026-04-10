package model

type Ataque struct {
	Nome   string `json:"nome_ataque"`
	Dano   int32  `json:"dano_ataque"`
	Custo  string `json:"custo_ataque"`
	Efeito string `json:"efeito_ataque"`
}
