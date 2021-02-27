package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "github.com/grpc-go-labs/demo1/solucion/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedElejirEquipoServer
}

type equipo struct {
	nombre pb.Equipo
	ciudad string
}

var equipos = []equipo{
	{
		nombre: pb.Equipo_AGUILAS,
		ciudad: "Santiago",
	},
	{
		nombre: pb.Equipo_LICEY,
		ciudad: "Santo Domingo",
	},
	{
		nombre: pb.Equipo_ESCOGIDO,
		ciudad: "Santo Domingo",
	},
	{
		nombre: pb.Equipo_TOROS,
		ciudad: "La Romana",
	},
	{
		nombre: pb.Equipo_ESTRELLAS,
		ciudad: "San Pedro de Macoris",
	},
	{
		nombre: pb.Equipo_GIGANTERS,
		ciudad: "San Francisco de Macoris",
	},
}

func (s *server) GetCiudad(ctx context.Context, in *pb.CiudadRequest) (*pb.CiudadRepuesta, error) {
	for _, e := range equipos {
		if e.nombre == in.Equipo {
			result := &pb.CiudadRepuesta{Ciudad: e.ciudad}
			return result, nil

		}
	}
	return nil, errors.New("Equipo no es valido")
}
func main() {
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
