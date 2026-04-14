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

func (r *Repository) GetDeck(id int32) (model.Response, error) {
	var response model.Response
	query := `SELECT id, quantidade, tipo, estatistica FROM deck WHERE id=$1`
	err := r.db.QueryRow(query, id).Scan(&response.Id, &response.Quantidade, &response.Tipo, &response.Estatistica.Id)
	if err != nil {
		fmt.Printf("Erro ao Buscar Deck : %v\n", err)
		return model.Response{}, err
	}

	query = `SELECT vitoria, derrota, total, pontos_ganho, pontos_perdido, media_pontos FROM estatistica WHERE id=$1`
	err = r.db.QueryRow(query, response.Estatistica.Id).Scan(&response.Estatistica.Vitoria, &response.Estatistica.Derrota, &response.Estatistica.Total, &response.Estatistica.Ponto_ganho, &response.Estatistica.Ponto_perdido, &response.Estatistica.Media_pontos)
	if err != nil {
		fmt.Printf("Erro ao Buscar Estatistica : %v\n", err)
		return model.Response{}, err
	}

	query = `SELECT card_deck FROM deck_card WHERE id_deck=$1`
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

	for _, crd := range listCard {
		var valueInt int32
		var cardUnit model.Card
		query = `SELECT id, type_card FROM cards`
		err := r.db.QueryRow(query, crd).Scan(&cardUnit.Id, &cardUnit.CardType)
		if err != nil {
			fmt.Printf("Erro ao Buscar card do Deck : %v\n", err)
			return model.Response{}, err
		}

		if cardUnit.CardType == "Pokemon" {
			query = `SELECT card_pokemon FROM pokemon_card WHERE id_card=$1`
			err := r.db.QueryRow(query, cardUnit.Id).Scan(&valueInt)
			if err != nil {
				fmt.Printf("Erro ao Buscar card do Deck : %v\n", err)
				return model.Response{}, err
			}
			cardUnit.Pokemon, err = r.GetCardPokemon(valueInt)
			if err != nil {
				fmt.Printf("Erro ao Buscar card do Deck : %v\n", err)
				return model.Response{}, err
			}
			response.Card = append(response.Card, cardUnit)
		} else if cardUnit.CardType == "Apoiador" {
			query = `SELECT card_apoiador FROM apoiador_card WHERE id_card=$1`
			err := r.db.QueryRow(query, cardUnit.Id).Scan(&valueInt)
			if err != nil {
				fmt.Printf("Erro ao Buscar card do Deck : %v\n", err)
				return model.Response{}, err
			}
			cardUnit.Apoiador, err = r.GetCardApoiador(valueInt)
			if err != nil {
				fmt.Printf("Erro ao Buscar card do Deck : %v\n", err)
				return model.Response{}, err
			}
			response.Card = append(response.Card, cardUnit)
		} else if cardUnit.CardType == "Item" {
			query = `SELECT card_item FROM item_card WHERE id_card=$1`
			err := r.db.QueryRow(query, cardUnit.Id).Scan(&valueInt)
			if err != nil {
				fmt.Printf("Erro ao Buscar card do Deck : %v\n", err)
				return model.Response{}, err
			}
			cardUnit.Item, err = r.GetCardItem(valueInt)
			if err != nil {
				fmt.Printf("Erro ao Buscar card do Deck : %v\n", err)
				return model.Response{}, err
			}
			response.Card = append(response.Card, cardUnit)
		}
	}
	return response, err
}

func (r *Repository) GetCardPokemon(id int32) (model.Pokemon, error) {
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

func (r *Repository) GetCardApoiador(id int32) (model.Apoiador, error) {
	var apoiador model.Apoiador
	query := `SELECT id, nome, card_type, efeito FROM apoiador WHERE id=$1`
	err := r.db.QueryRow(query, id).Scan(&apoiador.Id, &apoiador.Nome, &apoiador.CardType, &apoiador.Efeito)
	if err != nil {
		fmt.Printf("Erro ao Buscar card do Deck : %v\n", err)
		return model.Apoiador{}, err
	}
	return apoiador, err
}

func (r *Repository) GetCardItem(id int32) (model.Item, error) {
	var item model.Item
	query := `SELECT id, nome, card_type, efeito FROM apoiador WHERE id=$1`
	err := r.db.QueryRow(query, id).Scan(&item.Id, &item.Nome, &item.CardType, &item.Efeito)
	if err != nil {
		fmt.Printf("Erro ao Buscar card do Deck : %v\n", err)
		return model.Item{}, err
	}
	return item, err
}

func (r *Repository) CreateDeck(response model.Response) (int32, error) {
	var id int32
	var media float32
	total := response.Estatistica.Vitoria + response.Estatistica.Derrota
	if total != 0 {
		media = float32(response.Estatistica.Ponto_ganho / (response.Estatistica.Vitoria + response.Estatistica.Derrota))
	}
	query := `INSERT INTO estatistica (vitoria, derrota, total, pontos_ganho, pontos_perdido, media_pontos) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := r.db.QueryRow(query,
		response.Estatistica.Vitoria,
		response.Estatistica.Derrota,
		total,
		response.Estatistica.Ponto_ganho,
		response.Estatistica.Ponto_perdido,
		media).Scan(&response.Estatistica.Id)
	if err != nil {
		fmt.Printf("Erro ao criar a estatistica do Deck: %v\n", err)
		return 0, err
	}

	query = `INSERT INTO deck (quantidade, tipo, estatistica) VALUES ($1, $2, $3) RETURNING id`
	err = r.db.QueryRow(query, response.Quantidade, response.Tipo, response.Estatistica.Id).Scan(&id)
	if err != nil {
		fmt.Printf("Erro ao criar um Deck: %v\n", err)
		return 0, err
	}
	return id, err
}

func (r *Repository) AddCard(addCard model.AddCard) (model.Response, error) {
	for _, card := range addCard.Cards {
		var id int32
		query := `INSERT INTO cards (type_card) VALUES ($1) RETURNING id`
		err := r.db.QueryRow(query, card.CardType).Scan(&id)
		if err != nil {
			fmt.Printf("Erro ao adicionar um card no Deck: %v\n", err)
			return model.Response{}, err
		}
		query = `INSERT INTO deck_card (id_deck, card_deck) VALUES ($1, $2)`
		_, err = r.db.Exec(query, addCard.Id, id)
		if err != nil {
			fmt.Printf("Erro ao adicionar um card no Deck: %v\n", err)
			return model.Response{}, err
		}

		if card.CardType == "Pokemon" {
			query = `INSERT INTO cards_pokemon (id_card, card_pokemon) VALUES ($1, $2)`
			_, err = r.db.Exec(query, id, card.Pokemon.Id)
			if err != nil {
				fmt.Printf("Erro ao adicionar uma card Pokemon no Deck: %v\n", err)
				return model.Response{}, err
			}
		} else if card.CardType == "Apoiador" {
			query = `INSERT INTO cards_apoiador (id_card, card_apoiador) VALUES ($1, $2)`
			_, err = r.db.Exec(query, id, card.Apoiador.Id)
			if err != nil {
				fmt.Printf("Erro ao adicionar uma card Apoiador no Deck: %v\n", err)
				return model.Response{}, err
			}
		} else if card.CardType == "Item" {
			query = `INSERT INTO cards_item (id_card, card_item) VALUES ($1, $2)`
			_, err = r.db.Exec(query, id, card.Item.Id)
			if err != nil {
				fmt.Printf("Erro ao adicionar uma card Item no Deck: %v\n", err)
				return model.Response{}, err
			}
		}
	}
	response, err := r.GetDeck(addCard.Id)
	return response, err
}
