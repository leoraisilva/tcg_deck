package model

type Estatistica struct {
	Id            int32   `json:"id"`
	Vitoria       int32   `json:"vitoria"`
	Derrota       int32   `json:"derrota"`
	Total         int32   `json:"total"`
	Ponto_ganho   int32   `json:"pontos_ganho"`
	Ponto_perdido int32   `json:"pontos_perdido"`
	Media_pontos  float32 `json:"media_pontos"`
}
