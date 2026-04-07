package cmd

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pd "go/proto/deck_pocket_grpc.pd.go"
)

type server struct {
	pd.UnimplementedDeckServiceServer
}

func (s *server) CreateDeck(context.Context, *pd.Request) (*pd.Response, error) {
	return nil, status.Error(codes.Unimplemented, "method CreateDeck not implemented")
}
func (s *server) GetDeck(context.Context, *pd.GetDeckResquest) (*pd.Response, error) {
	return nil, status.Error(codes.Unimplemented, "method GetDeck not implemented")
}
