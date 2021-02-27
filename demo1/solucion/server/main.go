package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "google.golang.org/labs/demo1/solucion/proto"
)

type server struct {
	pb.UnimplementedElejirEquipoServer
}

func (s *server) GetCiudad(ctx context.Context, in *pb.CiudadRequest) (*pb.CiudadRepuesta, error) {
	ciudad := "Santiago"
	result := &pb.CiudadRepuesta{Ciudad: ciudad}
	return result, nil
}
func main() {
	fmt.Println("vim-go")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterElejirEquipoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
