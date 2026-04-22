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
	tx, err := r.db.Begin()
	if err != nil {
		fmt.Printf("Erro ao estabelecer conexao ao banco: %v\n", err)
		return model.Response{}, err
	}
	var response model.Response
	query := `SELECT id, quantidade, tipo, estatistica FROM deck WHERE id=$1`
	err = tx.QueryRow(query, id).Scan(&response.Id, &response.Quantidade, &response.Tipo, &response.Estatistica.Id)
	if err != nil {
		fmt.Printf("Erro no Repository ao Buscar Deck : %v\n", err)
		tx.Rollback()
		return model.Response{}, err
	}

	query = `SELECT vitoria, derrota, total, pontos_ganho, pontos_perdido, media_pontos FROM estatistica WHERE id=$1`
	err = tx.QueryRow(query, response.Estatistica.Id).Scan(&response.Estatistica.Vitoria, &response.Estatistica.Derrota, &response.Estatistica.Total, &response.Estatistica.Ponto_ganho, &response.Estatistica.Ponto_perdido, &response.Estatistica.Media_pontos)
	if err != nil {
		fmt.Printf("Erro no Repository ao Buscar Estatistica : %v\n", err)
		tx.Rollback()
		return model.Response{}, err
	}

	query = `SELECT card_deck FROM deck_card WHERE id_deck=$1`
	card, err := tx.Query(query, id)
	if err != nil {
		fmt.Printf("Erro no Repository ao Buscar card do Deck : %v\n", err)
		tx.Rollback()
		return model.Response{}, err
	}
	var listCard []int32
	for card.Next() {
		var cardUnit int32
		if err = card.Scan(&cardUnit); err != nil {
			fmt.Printf("Erro no Repository ao Buscar card do Deck : %v\n", err)
			tx.Rollback()
			return model.Response{}, err
		}
		listCard = append(listCard, cardUnit)
	}

	for _, crd := range listCard {
		var valueInt int32
		var cardUnit model.CardResponse
		cardUnit.Id = crd
		query = `SELECT type_card FROM cards WHERE id=$1`
		err := tx.QueryRow(query, cardUnit.Id).Scan(&cardUnit.CardType)
		if err != nil {
			fmt.Printf("Erro no Repository ao Buscar card do Deck : %v\n", err)
			tx.Rollback()
			return model.Response{}, err
		}

		switch cardUnit.CardType {
		case "Pokemon":
			query = `SELECT card_pokemon FROM cards_pokemon WHERE id_card=$1`
			err := tx.QueryRow(query, cardUnit.Id).Scan(&valueInt)
			if err != nil {
				fmt.Printf("Erro no Repository ao Buscar card do Deck : %v\n", err)
				tx.Rollback()
				return model.Response{}, err
			}
			cardUnit.Pokemon, err = r.GetCardPokemon(tx, valueInt)
			if err != nil {
				fmt.Printf("Erro no Repository ao Buscar card do Deck : %v\n", err)
				tx.Rollback()
				return model.Response{}, err
			}
			response.Card = append(response.Card, cardUnit)
		case "Apoiador":
			query = `SELECT card_apoiador FROM cards_apoiador WHERE id_card=$1`
			err := tx.QueryRow(query, cardUnit.Id).Scan(&valueInt)
			if err != nil {
				fmt.Printf("Erro no Repository ao Buscar card do Deck : %v\n", err)
				tx.Rollback()
				return model.Response{}, err
			}
			cardUnit.Apoiador, err = r.GetCardApoiador(tx, valueInt)
			if err != nil {
				fmt.Printf("Erro no Repository ao Buscar card do Deck : %v\n", err)
				tx.Rollback()
				return model.Response{}, err
			}
			response.Card = append(response.Card, cardUnit)
		case "Item":
			query = `SELECT card_item FROM cards_item WHERE id_card=$1`
			err := tx.QueryRow(query, cardUnit.Id).Scan(&valueInt)
			if err != nil {
				fmt.Printf("Erro ao Buscar card do Deck : %v\n", err)
				tx.Rollback()
				return model.Response{}, err
			}
			cardUnit.Item, err = r.GetCardItem(tx, valueInt)
			if err != nil {
				fmt.Printf("Erro no Repository ao Buscar card do Deck : %v\n", err)
				tx.Rollback()
				return model.Response{}, err
			}
			response.Card = append(response.Card, cardUnit)
		}
	}
	return response, tx.Commit()
}

func (r *Repository) GetCardPokemon(tx *sql.Tx, id int32) (model.Pokemon, error) {
	var pokemon model.Pokemon
	query := `SELECT id, nome, card_type, tipo, estagio, geracao, ps, recuo, fraqueza FROM pokemon WHERE id=$1`
	err := tx.QueryRow(query, id).Scan(&pokemon.Id, &pokemon.Nome, &pokemon.TipoCarta, &pokemon.Tipo, &pokemon.Estagio, &pokemon.Geracao, &pokemon.PS, &pokemon.Recuo, &pokemon.Fraqueza)
	if err != nil {
		fmt.Printf("Erro ao Buscar card do Deck : %v\n", err)
		tx.Rollback()
		return model.Pokemon{}, err
	}
	query = `SELECT a.nome_ataque, a.dano_ataque, a.custo_ataque, a.efeito_ataque
	FROM ataque a
	JOIN pokemon_ataque pa ON pa.ataque = a.nome_ataque
	WHERE pa.id_pokemon = $1`
	ataque, err := tx.Query(query, id)
	if err != nil {
		fmt.Printf("Erro ao Buscar ataque do pokemon: %v\n", err)
		tx.Rollback()
		return model.Pokemon{}, err
	}
	defer ataque.Close()

	for ataque.Next() {
		var valueAtaque model.Ataque
		if err = ataque.Scan(&valueAtaque.Nome, &valueAtaque.Dano, &valueAtaque.Custo, &valueAtaque.Efeito); err != nil {
			fmt.Printf("Erro ao Buscar ataque do pokemon: %v\n", err)
			tx.Rollback()
			return model.Pokemon{}, err
		}
		pokemon.Ataque = append(pokemon.Ataque, valueAtaque)
	}

	query = `SELECT h.nome_habilidade, h.efeito_habilidade
	FROM habilidade h
	JOIN pokemon_habilidade ph ON ph.habilidade = h.nome_habilidade
	WHERE ph.id_pokemon = $1`
	habilidade, err := tx.Query(query, id)
	if err != nil {
		fmt.Printf("Erro ao Buscar habilidade do pokemon: %v\n", err)
		tx.Rollback()
		return model.Pokemon{}, err
	}

	habilidade.Close()

	for habilidade.Next() {
		var valueHabilidade model.Habilidade
		if err = habilidade.Scan(&valueHabilidade.Nome, &valueHabilidade.Efeito); err != nil {
			fmt.Printf("Erro ao Buscar Habilidade do pokemon: %v\n", err)
			tx.Rollback()
			return model.Pokemon{}, err
		}
		pokemon.Habilidade = append(pokemon.Habilidade, valueHabilidade)
	}

	return pokemon, err
}

func (r *Repository) GetCardApoiador(tx *sql.Tx, id int32) (model.Apoiador, error) {
	var apoiador model.Apoiador
	query := `SELECT id, nome, card_type, efeito FROM apoiador WHERE id=$1`
	err := tx.QueryRow(query, id).Scan(&apoiador.Id, &apoiador.Nome, &apoiador.CardType, &apoiador.Efeito)
	if err != nil {
		fmt.Printf("Erro ao Buscar card do Deck : %v\n", err)
		tx.Rollback()
		return model.Apoiador{}, err
	}
	return apoiador, err
}

func (r *Repository) GetCardItem(tx *sql.Tx, id int32) (model.Item, error) {
	var item model.Item
	query := `SELECT id, nome, card_type, efeito FROM apoiador WHERE id=$1`
	err := tx.QueryRow(query, id).Scan(&item.Id, &item.Nome, &item.CardType, &item.Efeito)
	if err != nil {
		fmt.Printf("Erro ao Buscar card do Deck : %v\n", err)
		tx.Rollback()
		return model.Item{}, err
	}
	return item, err
}

func (r *Repository) CreateDeck(response model.Response) (int32, error) {
	var id int32
	var media float32
	tx, err := r.db.Begin()
	if err != nil {
		fmt.Printf("Erro ao estabelecer conexao ao banco: %v\n", err)
		return 0, err
	}
	total := response.Estatistica.Vitoria + response.Estatistica.Derrota
	if total != 0 {
		media = float32(response.Estatistica.Ponto_ganho) / float32(total)
	} else {
		media = 0
	}
	query := `INSERT INTO estatistica (vitoria, derrota, total, pontos_ganho, pontos_perdido, media_pontos) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err = tx.QueryRow(query,
		response.Estatistica.Vitoria,
		response.Estatistica.Derrota,
		total,
		response.Estatistica.Ponto_ganho,
		response.Estatistica.Ponto_perdido,
		media).Scan(&response.Estatistica.Id)
	if err != nil {
		fmt.Printf("Erro ao criar a estatistica do Deck: %v\n", err)
		tx.Rollback()
		return 0, err
	}

	query = `INSERT INTO deck (quantidade, tipo, estatistica) VALUES ($1, $2, $3) RETURNING id`
	err = tx.QueryRow(query, response.Quantidade, response.Tipo, response.Estatistica.Id).Scan(&id)
	if err != nil {
		fmt.Printf("Erro ao criar um Deck: %v\n", err)
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *Repository) AddCard(addCard model.AddCard) (model.Response, error) {
	tx, err := r.db.Begin()
	if err != nil {
		fmt.Printf("Erro ao estabelecer conexao ao banco: %v\n", err)
		return model.Response{}, err
	}
	for _, card := range addCard.Cards {
		var id int32
		query := `INSERT INTO cards (type_card) VALUES ($1) RETURNING id`
		err := tx.QueryRow(query, card.CardType).Scan(&id)
		if err != nil {
			fmt.Printf("Erro ao adicionar um card no Deck: %v\n", err)
			tx.Rollback()
			return model.Response{}, err
		}
		query = `INSERT INTO deck_card (id_deck, card_deck) VALUES ($1, $2)`
		_, err = tx.Exec(query, addCard.Id, id)
		if err != nil {
			fmt.Printf("Erro ao adicionar um card no Deck: %v\n", err)
			tx.Rollback()
			return model.Response{}, err
		}

		switch card.CardType {
		case "Pokemon":
			query = `INSERT INTO cards_pokemon (id_card, card_pokemon) VALUES ($1, $2)`
			_, err = tx.Exec(query, id, card.Pokemon)
			if err != nil {
				fmt.Printf("Erro ao adicionar uma card Pokemon no Deck: %v\n", err)
				tx.Rollback()
				return model.Response{}, err
			}
		case "Apoiador":
			query = `INSERT INTO cards_apoiador (id_card, card_apoiador) VALUES ($1, $2)`
			_, err = tx.Exec(query, id, card.Apoiador)
			if err != nil {
				fmt.Printf("Erro ao adicionar uma card Apoiador no Deck: %v\n", err)
				tx.Rollback()
				return model.Response{}, err
			}
		case "Item":
			query = `INSERT INTO cards_item (id_card, card_item) VALUES ($1, $2)`
			_, err = tx.Exec(query, id, card.Item)
			if err != nil {
				fmt.Printf("Erro ao adicionar uma card Item no Deck: %v\n", err)
				tx.Rollback()
				return model.Response{}, err
			}
		}
	}
	response, err := r.GetDeck(addCard.Id)
	return response, tx.Commit()
}

func (r *Repository) EditDeck(editDeck model.EditCard) (model.Response, error) {
	tx, err := r.db.Begin()
	if err != nil {
		fmt.Printf("Erro ao estabelecer conexao ao banco: %v\n", err)
		return model.Response{}, err
	}
	for _, card := range editDeck.Cards {
		query := `UPDATE cards SET type_card=$1 WHERE id=$2`
		_, err := tx.Exec(query, card.CardType, card.Id)
		if err != nil {
			fmt.Printf("Erro ao alterar cards no deck: %v\n", err)
			tx.Rollback()
			return model.Response{}, err
		}
		switch card.CardType {
		case "Pokemon":
			query = `UPDATE cards_pokemon SET card_pokemon=$1 WHERE id_card=$2`
			_, err = tx.Exec(query, card.Pokemon, card.Id)
			if err != nil {
				fmt.Printf("Erro ao alterar cards pokemon no deck: %v\n", err)
				tx.Rollback()
				return model.Response{}, err
			}
		case "Apoiador":
			query = `UPDATE cards_apoiador SET card_apoiador=$1 WHERE id_card=$2`
			_, err = tx.Exec(query, card.Apoiador, card.Id)
			if err != nil {
				fmt.Printf("Erro ao alterar cards apoiador no deck: %v\n", err)
				tx.Rollback()
				return model.Response{}, err
			}
		case "Item":
			query = `UPDATE cards_item SET card_item=$1 WHERE id_card=$2`
			_, err = tx.Exec(query, card.Item, card.Id)
			if err != nil {
				fmt.Printf("Erro ao alterar cards item no deck: %v\n", err)
				tx.Rollback()
				return model.Response{}, err
			}
		}
	}
	response, err := r.GetDeck(editDeck.Id)
	if err != nil {
		fmt.Printf("Erro ao Buscar o deck: %v\n", err)
		tx.Rollback()
		return model.Response{}, err
	}
	return response, tx.Commit()
}

func (r *Repository) RemoveDeck(id int32) (model.ResponseMessage, error) {
	var estatistica int32
	tx, err := r.db.Begin()
	if err != nil {
		fmt.Printf("Erro ao estabelecer conexao ao banco: %v\n", err)
		return model.ResponseMessage{}, err
	}
	query := `SELECT estatistica FROM deck WHERE id=$1`
	err = tx.QueryRow(query, id).Scan(&estatistica)
	if err != nil {
		fmt.Printf("Erro ao localizar estatistica: %v\n", err)
		tx.Rollback()
		return model.ResponseMessage{}, err
	}
	query = `DELETE FROM estatistica WHERE id=$1`
	_, err = tx.Exec(query, estatistica)
	if err != nil {
		fmt.Printf("Erro ao deletar estatistica: %v\n", err)
		tx.Rollback()
		return model.ResponseMessage{}, err
	}
	query = `DELETE FROM deck_card WHERE id_deck=$1 CASCADE`
	_, err = tx.Exec(query, id)
	if err != nil {
		fmt.Printf("Erro ao deletar deck card: %v\n", err)
		tx.Rollback()
		return model.ResponseMessage{}, err
	}
	query = `DELETE FROM deck WHERE id=$1`
	_, err = tx.Exec(query, id)
	if err != nil {
		fmt.Printf("Erro ao deletar deck: %v\n", err)
		tx.Rollback()
		return model.ResponseMessage{}, err
	}
	var message model.ResponseMessage
	message.Message = "Deletado com sucesso !!"
	return message, tx.Commit()
}
