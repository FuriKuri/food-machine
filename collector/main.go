//go:generate protoc -I ../fruit --go_out=plugins=grpc:../fruit ../fruit/fruit.proto

package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/furikuri/fruit-generator/fruit"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}


func (s *server) Deliver(ctx context.Context, in *pb.Fruit) (*pb.Empty, error) {
	log.Print(in.Name)
	return &pb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFruitCollectorServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
