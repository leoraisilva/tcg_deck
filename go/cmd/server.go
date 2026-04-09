package cmd

import (
	"context"
	"tcg_deck/go/model"
	"tcg_deck/go/usecase"

	pd "go/proto/deck_pocket_grpc.pd.go"
)

type server struct {
	pd.UnimplementedDeckServiceServer
	usecase.Usecase
}

func (s *server) CreateDeck(cont context.Context, request *pd.Request) *pd.Response {
	var response model.Response
	response.Quantidade = request.GetQuantidade()
	response.Tipo = request.GetTipo()
	response.Card = request.GetCard()
	response.Estatistica = request.GetEstatistica()
	s.Usecase.CreateDeck(response)

	&pd.CreateDeck(request)
	return response
}
func (s *server) GetDeck(cont context.Context, request *pd.GetDeckResquest) (*pd.Response, error) {
	&pd.GetDeck(request)
	return s.Usecase.GetDeck(request.GetID)
}
