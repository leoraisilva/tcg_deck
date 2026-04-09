package cmd

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("erro ao abrir porta: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{})

	log.Println("gRPC server rodando em :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("erro no servidor: %v", err)
	}
}
