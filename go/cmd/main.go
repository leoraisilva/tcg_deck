package main

import (
	"context"
	"log"
	"net"
	"tcg_deck/go/helper"
	"tcg_deck/go/mapper"
	"tcg_deck/go/pb"
	"tcg_deck/go/repository"
	"tcg_deck/go/usecase"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedDeckServiceServer
	usecase usecase.Usecase
	mapper  mapper.Mapper
}

func (s *Server) CreateDeck(cont context.Context, request *pb.Request) (*pb.Response, error) {
	response := s.mapper.ToModel(request)
	_, err := s.usecase.CreateDeck(response)
	if err != nil {
		panic(err)
	}

	result := s.mapper.ToPBResponse(response)
	return result, nil
}

func (s *Server) GetDeck(cont context.Context, request *pb.GetDeckResquest) (*pb.Response, error) {
	response, err := s.usecase.GetDeck(request.GetID())
	if err != nil {
		panic(err)
	}

	result := s.mapper.ToPBResponse(response)
	return result, nil
}

func (s *Server) AddCard(cont context.Context, request *pb.AddCardRequest) (*pb.Response, error) {
	response, err := s.usecase.AddCard(s.mapper.ToAddCardModel(request))
	if err != nil {
		panic(err)
	}
	return s.mapper.ToPBResponse(response), nil
}

func (s *Server) EditDeck(cont context.Context, req *pb.EditCardRequest) (*pb.Response, error) {
	response, err := s.usecase.EditDeck(s.mapper.ToEditModel(req))
	if err != nil {
		panic(err)
	}
	return s.mapper.ToPBResponse(response), nil
}

func (s *Server) RemoveDeck(cont context.Context, req *pb.GetDeckResquest) (*pb.ResponseMessage, error) {
	response, err := s.usecase.RemoveDeck(req.ID)
	if err != nil {
		panic(err)
	}
	var result *pb.ResponseMessage
	result.Response = response.Message
	return result, nil
}

func main() {

	conn, err := helper.GetConnection()
	if err != nil {
		panic(err)
	}
	repository := repository.NewRepository(conn)
	usecase := usecase.NewUsecase(repository)
	mapper := mapper.NewMapper()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("erro ao abrir porta: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDeckServiceServer(grpcServer, &Server{usecase: usecase, mapper: mapper})

	log.Println("gRPC server rodando em :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("erro no servidor: %v", err)
	}
}
