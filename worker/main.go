package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"

	pb "github.com/furikuri/food-machine/food"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

var machineID = pseudoUUID()

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFoodCollectorClient(conn)

	ticker := time.NewTicker(700 * time.Millisecond)
	for t := range ticker.C {
		fmt.Println("Tick at", t)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		_, err = c.Deliver(ctx, &pb.Food{Name: pick(), Worker: machineID})
		if err != nil {
			log.Fatalf("could not food: %v", err)
		}
	}
	log.Printf("FINISHED")
}

func pseudoUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return ""
	}

	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
