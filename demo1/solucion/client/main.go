package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/amb1s1/grpc-go-labs/demo1/solucion/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect to %v, error: %v", address, err)
	}
	defer conn.Close()

	c := pb.NewElejirEquipoClient(conn)

	ctx := context.Background()
	request := &pb.CiudadRequest{
		Equipo: pb.Equipo_LICEY,
	}
	response, err := c.GetCiudad(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("la ciudad es: %v\n", response.Ciudad)
}
