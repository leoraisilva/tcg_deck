package model

type Estatistica struct {
	Id            int     `json:"id"`
	Vitoria       int     `json:"vitoria"`
	Derrota       int     `json:"derrota"`
	Total         int     `json:"total"`
	Ponto_ganho   int     `json:"pontos_ganho"`
	Ponto_perdido int     `json:"pontos_perdido"`
	Media_pontos  float32 `json:"media_pontos"`
}
