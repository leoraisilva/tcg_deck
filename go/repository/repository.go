package repository

import (
	"database/sql"
	"fmt"
	"tcg_deck/go/model"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{db: db}
}

func (r *Repository) GetDeck(id int) (model.Response, error) {
	var response model.Response
	var estatistica int
	query := `SELECT id, quantidade, tipo, estatistica FROM deck_tb WHERE id=$1`
	err := r.db.QueryRow(query, id).Scan(&response.Id, &response.Quantidade, &response.Tipo, &estatistica)
	if err != nil {
		fmt.Printf("Erro ao Buscar Deck : %v\n", err)
		return model.Response{}, err
	}

	query = `SELECT id, vitoria, derrota, total, pontos_ganho, pontos_perdido, media_pontos FROM estatistica WHERE id=$1`
	err = r.db.QueryRow(query, id).Scan(&response.Estatistica.Id, &response.Estatistica.Vitoria, &response.Estatistica.Derrota, &response.Estatistica.Total, &response.Estatistica.Ponto_ganho, &response.Estatistica.Ponto_perdido, &response.Estatistica.Media_pontos)
	if err != nil {
		fmt.Printf("Erro ao Buscar Estatistica : %v\n", err)
		return model.Response{}, err
	}

	query = `SELECT id_card FROM deck_card WHERE id_deck=$1`
	card, err := r.db.Query(query, id)
	if err != nil {
		fmt.Printf("Erro ao Buscar card do Deck : %v\n", err)
		return model.Response{}, err
	}
	for _, cardDeck := range card {

	}

}
