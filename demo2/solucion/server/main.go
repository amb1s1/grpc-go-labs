package main

import (
	"log"
	"net"

	pb "github.com/amb1s1/grpc-go-labs/demo2/solucion/proto"
)

type server struct {
	pb.UnimplementedElejirEquipoServer
}

func main() {
	_, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to start listening: %v\n", err)

	}
}
