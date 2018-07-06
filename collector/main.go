//go:generate protoc -I ../fruit --go_out=plugins=grpc:../fruit ../fruit/fruit.proto

package main

import (
	"net"

	"net/http"

	pb "github.com/furikuri/fruit-generator/fruit"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	
	"github.com/gorilla/websocket"
	"log"
)

const (
	port = ":50051"
)

type server struct{}

var listenerSockets = make(map[chan string]bool)

func (s *server) Deliver(ctx context.Context, in *pb.Fruit) (*pb.Empty, error) {
	log.Print(in.Name)
	for listener, _ := range listenerSockets { 
	    listener <- in.Name
	}
	return &pb.Empty{}, nil
}

var upgrader = websocket.Upgrader{} // use default options

func fruits(w http.ResponseWriter, r *http.Request) {
	channel := make(chan string)
	listenerSockets[channel] = true
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	defer delete(listenerSockets, channel)
	for {
		fruit := <-channel
		log.Printf("recv: %s", fruit)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/fruits", fruits)
	go http.ListenAndServe(":8080", nil)

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
