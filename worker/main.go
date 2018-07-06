package main

import (
	"log"
	"time"

	pb "github.com/furikuri/fruit-generator/fruit"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFruitCollectorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.Deliver(ctx, &pb.Fruit{Name: "apple", Worker: "1"})
	if err != nil {
		log.Fatalf("could not fruit: %v", err)
	}
	log.Printf("OK")
}
