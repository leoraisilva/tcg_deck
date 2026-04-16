package mapper

import (
	"tcg_deck/go/model"
	"tcg_deck/go/pb"
)

type Mapper struct{}

func NewMapper() Mapper {
	return Mapper{}
}

func (m *Mapper) ToModel(request *pb.Request) model.Response {
	var response model.Response
	response.Quantidade = request.Quantidade
	response.Tipo = model.Tipo(request.Tipo)
	response.Estatistica = toEstatisticaModel(request.Estatistica)
	return response
}

func (m *Mapper) ToAddCardModel(req *pb.AddCardRequest) model.AddCard {
	var addCard model.AddCard
	addCard.Id = req.IdDeck
	for _, card := range req.Card {
		var cardUnit model.Card
		cardUnit.Id = card.ID
		cardUnit.CardType = card.CardType
		switch cardUnit.CardType {
		case "Pokemon":
			cardUnit.Pokemon = card.IdPokemon
		case "Apoiador":
			cardUnit.Apoiador = card.IdApoiador
		case "Item":
			cardUnit.Item = card.IdItem
		}
		addCard.Cards = append(addCard.Cards, cardUnit)
	}
	return addCard
}

func (m *Mapper) ToAddCardRequest(addCard model.AddCard) *pb.AddCardRequest {
	var addCardRequest pb.AddCardRequest
	addCardRequest.IdDeck = addCard.Id
	for _, card := range addCard.Cards {
		var cardUnit pb.Card
		cardUnit.ID = card.Id
		cardUnit.CardType = card.CardType
		switch cardUnit.CardType {
		case "Pokemon":
			cardUnit.IdPokemon = card.Pokemon
		case "Apoiador":
			cardUnit.IdApoiador = card.Apoiador
		case "Item":
			cardUnit.IdItem = card.Item
		}
		addCardRequest.Card = append(addCardRequest.Card, &cardUnit)
	}
	return &addCardRequest
}

func (m *Mapper) ToPBResponse(response model.Response) *pb.Response {
	var pbResp pb.Response

	pbResp.Quantidade = response.Quantidade
	pbResp.Tipo = string(response.Tipo)
	pbResp.Estatistica = toEstatisticaPB(response.Estatistica)

	for _, card := range response.Card {
		var cardUnit pb.CardResponse
		cardUnit.ID = card.Id
		cardUnit.CardType = card.CardType
		switch cardUnit.CardType {
		case "Pokemon":
			cardUnit.Pokemon = toPokemonPB(card.Pokemon)
		case "Apoiador":
			cardUnit.Apoiador = toApoiadorPB(card.Apoiador)
		case "Item":
			cardUnit.Item = toItemPB(card.Item)
		}
		pbResp.Card = append(pbResp.Card, &cardUnit)
	}
	return &pbResp
}

func (m *Mapper) ToPBEditRequest(editDeck model.EditCard) *pb.EditCardRequest {
	var editRequest pb.EditCardRequest
	editRequest.IdDeck = editDeck.Id
	for _, card := range editDeck.Cards {
		var cardRequest pb.Card
		cardRequest.ID = card.Id
		cardRequest.CardType = card.CardType
		switch cardRequest.CardType {
		case "Pokemon":
			cardRequest.IdPokemon = card.Pokemon
		case "Apoiador":
			cardRequest.IdApoiador = card.Apoiador
		case "Item":
			cardRequest.IdItem = card.Item
		}
		editRequest.Card = append(editRequest.Card, &cardRequest)
	}
	return &editRequest
}

func (m *Mapper) ToEditModel(req *pb.EditCardRequest) model.EditCard {
	var editDeck model.EditCard
	editDeck.Id = req.IdDeck
	for _, card := range req.Card {
		var cardModel model.Card
		cardModel.Id = card.ID
		cardModel.CardType = card.CardType
		switch cardModel.CardType {
		case "Pokemon":
			cardModel.Pokemon = card.IdPokemon
		case "Apoiador":
			cardModel.Apoiador = card.IdApoiador
		case "Item":
			cardModel.Item = card.IdItem
		}
		editDeck.Cards = append(editDeck.Cards, cardModel)
	}
	return editDeck
}

func toPokemonModel(req *pb.Pokemon) model.Pokemon {
	var pokemon model.Pokemon
	pokemon.Id = req.ID
	pokemon.Nome = req.Nome
	pokemon.Tipo = model.Tipo(req.Tipo)
	pokemon.Estagio = req.Estagio
	pokemon.Geracao = req.Geracao
	pokemon.Recuo = req.Recuo
	pokemon.PS = req.Ps
	pokemon.Fraqueza = model.Tipo(req.Fraqueza)

	return pokemon
}

func toAtaqueModel(req *pb.Ataque) model.Ataque {
	var ataque model.Ataque
	ataque.Nome = req.Nome
	ataque.Dano = req.Dano
	ataque.Custo = req.Custo
	ataque.Efeito = req.Efeito

	return ataque
}

func toHabilidadeModel(req *pb.Habilidade) model.Habilidade {
	var habilidade model.Habilidade
	habilidade.Nome = req.Nome
	habilidade.Efeito = req.Efeito

	return habilidade
}

func toApoiadorModel(req *pb.Apoiador) model.Apoiador {
	var apoiador model.Apoiador
	apoiador.Id = req.ID
	apoiador.Nome = req.Nome
	apoiador.CardType = req.CardType
	apoiador.Efeito = req.Efeito

	return apoiador
}

func toItemModel(req *pb.Item) model.Item {
	var item model.Item
	item.Id = req.ID
	item.Nome = req.Nome
	item.CardType = req.CardType
	item.Efeito = append(item.Efeito, req.Efeito)

	return item
}

func toEstatisticaPB(model model.Estatistica) *pb.Estatistica {
	return &pb.Estatistica{
		Vitoria:      model.Vitoria,
		Derrota:      model.Derrota,
		PontoGanho:   model.Ponto_ganho,
		PontoPerdido: model.Ponto_perdido,
	}
}

func toPokemonPB(model model.Pokemon) *pb.Pokemon {
	return &pb.Pokemon{
		ID:       model.Id,
		Nome:     model.Nome,
		Tipo:     string(model.Tipo),
		Estagio:  model.Estagio,
		Geracao:  model.Geracao,
		Recuo:    model.Recuo,
		Ps:       model.PS,
		Fraqueza: string(model.Fraqueza),
	}
}

func toAtaquePB(model model.Ataque) *pb.Ataque {
	return &pb.Ataque{
		Nome:   model.Nome,
		Dano:   model.Dano,
		Custo:  model.Custo,
		Efeito: model.Efeito,
	}
}

func toHabilidadePB(model model.Habilidade) *pb.Habilidade {
	return &pb.Habilidade{
		Nome:   model.Nome,
		Efeito: model.Efeito,
	}
}

func toApoiadorPB(model model.Apoiador) *pb.Apoiador {
	return &pb.Apoiador{
		ID:       model.Id,
		Nome:     model.Nome,
		CardType: model.CardType,
		Efeito:   model.Efeito,
	}
}

func toItemPB(model model.Item) *pb.Item {
	var efeito string
	if len(model.Efeito) > 0 {
		efeito = model.Efeito[0]
	}

	return &pb.Item{
		ID:       model.Id,
		Nome:     model.Nome,
		CardType: model.CardType,
		Efeito:   efeito,
	}
}

func toEstatisticaModel(req *pb.Estatistica) model.Estatistica {
	var estatistica model.Estatistica
	estatistica.Vitoria = req.Vitoria
	estatistica.Derrota = req.Derrota
	estatistica.Total = req.Vitoria + req.Derrota
	estatistica.Ponto_ganho = req.PontoGanho
	estatistica.Ponto_perdido = req.PontoPerdido
	estatistica.Media_pontos = float32(req.PontoGanho / estatistica.Total)

	return estatistica
}
