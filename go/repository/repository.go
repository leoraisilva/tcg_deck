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
	var listCard []string
	for card.Next() {
		var cardUnit string
		if err = card.Scan(&cardUnit); err != nil {
			fmt.Printf("Erro ao Buscar card do Deck : %v\n", err)
			return model.Response{}, err
		}
		listCard = append(listCard, cardUnit)
	}

	for _, crd = range listCard {

	}

}

func (r *Repository) GetCardPokemon(id int) (model.Pokemon, error) {
	var pokemon model.Pokemon
	query := `SELECT id, nome, card_type, tipo, estagio, geracao, ps, recuo, fraqueza FROM pokemon WHERE id=$1`
	err := r.db.QueryRow(query, id).Scan(&pokemon.Id, &pokemon.Nome, &pokemon.TipoCarta, &pokemon.Tipo, &pokemon.Estagio, &pokemon.Geracao, &pokemon.PS, &pokemon.Recuo, &pokemon.Fraqueza)
	if err != nil {
		fmt.Printf("Erro ao Buscar card do Deck : %v\n", err)
		return model.Pokemon{}, err
	}
	query = `SELECT ataque FROM pokemon_ataque WHERE id_pokemon=$1`
	ataque, err := r.db.Query(query, id)
	if err != nil {
		fmt.Printf("Erro ao Buscar ataque do pokemon: %v\n", err)
		return model.Pokemon{}, err
	}
	var listAtaque []string
	for ataque.Next() {
		var ataqueUnit string
		if err = ataque.Scan(&ataqueUnit); err != nil {
			fmt.Printf("Erro ao Buscar ataque do pokemon: %v\n", err)
			return model.Pokemon{}, err
		}
		listAtaque = append(listAtaque, ataqueUnit)
	}

	for _, atk := range listAtaque {
		var valueAtaque model.Ataque
		query = `SELECT nome_ataque, dano_ataque, custo_ataque, efeito_ataque FROM ataque WHERE nome_ataque=$1`
		err = r.db.QueryRow(query, atk).Scan(&valueAtaque.Nome, &valueAtaque.Dano, &valueAtaque.Custo, &valueAtaque.Efeito)
		if err != nil {
			fmt.Printf("Erro ao Buscar ataque do pokemon: %v\n", err)
			return model.Pokemon{}, err
		}
		pokemon.Ataque = append(pokemon.Ataque, valueAtaque)
	}

	query = `SELECT habilidade FROM pokemon_habilidade WHERE id_pokemon=$1`
	habilidade, err := r.db.Query(query, id)
	if err != nil {
		fmt.Printf("Erro ao Buscar habilidade do pokemon: %v\n", err)
		return model.Pokemon{}, err
	}
	var listHabilidade []string
	for habilidade.Next() {
		var habilidadeUnit string
		if err = habilidade.Scan(&habilidadeUnit); err != nil {
			fmt.Printf("Erro ao Buscar habilidade do pokemon: %v\n", err)
			return model.Pokemon{}, err
		}
		listHabilidade = append(listHabilidade, habilidadeUnit)
	}

	for _, atk := range listHabilidade {
		var valueHabilidade model.Habilidade
		query = `SELECT nome_habilidade, efeito_habilidade FROM habilidade WHERE nome_habilidade=$1`
		err = r.db.QueryRow(query, atk).Scan(&valueHabilidade.Nome, &valueHabilidade.Efeito)
		if err != nil {
			fmt.Printf("Erro ao Buscar Habilidade do pokemon: %v\n", err)
			return model.Pokemon{}, err
		}
		pokemon.Habilidade = append(pokemon.Habilidade, valueHabilidade)
	}

	return pokemon, err
}

func (r *Repository) GetCardApoiador(id int) (model.Apoiador, error) {
	var apoiador model.Apoiador
	query := `SELECT id, nome, card_type, efeito FROM apoiador WHERE id=$1`
	err := r.db.QueryRow(query, id).Scan(&apoiador.Id, &apoiador.Nome, &apoiador.CardType, &apoiador.Efeito)
	if err != nil {
		fmt.Printf("Erro ao Buscar card do Deck : %v\n", err)
		return model.Apoiador{}, err
	}
	return apoiador, err
}

func (r *Repository) GetCardItem(id int) (model.Item, error) {
	var item model.Item
	query := `SELECT id, nome, card_type, efeito FROM apoiador WHERE id=$1`
	err := r.db.QueryRow(query, id).Scan(&item.Id, &item.Nome, &item.CardType, &item.Efeito)
	if err != nil {
		fmt.Printf("Erro ao Buscar card do Deck : %v\n", err)
		return model.Item{}, err
	}
	return item, err
}
