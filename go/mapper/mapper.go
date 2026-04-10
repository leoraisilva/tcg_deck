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
	response.Tipo = response.Tipo
	response.Estatistica = toEstatisticaModel(request.Estatistica)
	for _, card := range request.Card {
		var cardUnit model.Card
		cardUnit.Id = card.ID

		switch c := card.CardType.(type) {

		case *pb.Card_Pokemon:
			cardUnit.CardType = "Pokemon"
			cardUnit.Pokemon = toPokemonModel(c.Pokemon)

		case *pb.Card_Apoiador:
			cardUnit.CardType = "Apoiador"
			cardUnit.Apoiador = toApoiadorModel(c.Apoiador)

		case *pb.Card_Item:
			cardUnit.CardType = "Item"
			cardUnit.Item = toItemModel(c.Item)
		}

		response.Card = append(response.Card, cardUnit)
	}
	return response
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

func (m *Mapper) ToPBResponse(response model.Response) *pb.Response {
	var pbResp pb.Response

	pbResp.Quantidade = response.Quantidade
	pbResp.Tipo = string(response.Tipo)
	pbResp.Estatistica = toEstatisticaPB(response.Estatistica)

	for _, card := range response.Card {
		var pbCard pb.Card
		pbCard.ID = card.Id

		switch card.CardType {

		case "Pokemon":
			pbCard.CardType = &pb.Card_Pokemon{
				Pokemon: toPokemonPB(card.Pokemon),
			}

		case "Apoiador":
			pbCard.CardType = &pb.Card_Apoiador{
				Apoiador: toApoiadorPB(card.Apoiador),
			}

		case "Item":
			pbCard.CardType = &pb.Card_Item{
				Item: toItemPB(card.Item),
			}
		}

		pbResp.Card = append(pbResp.Card, &pbCard)
	}

	return &pbResp
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
