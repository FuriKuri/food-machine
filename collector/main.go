//go:generate protoc -I ../food --go_out=plugins=grpc:../food ../food/food.proto

package main

import (
	"net"
	"time"

	"net/http"

	pb "github.com/furikuri/food-machine/food"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"log"
	"sync/atomic"

	"github.com/gorilla/websocket"
)

const (
	port = ":50051"
)

type server struct{}

var listenerSockets = make(map[chan string]bool)
var ops uint64

func (s *server) Deliver(ctx context.Context, in *pb.Food) (*pb.Empty, error) {
	log.Print(in.Name)
	atomic.AddUint64(&ops, 1)
	for listener := range listenerSockets {
		listener <- in.Name
	}
	return &pb.Empty{}, nil
}

var upgrader = websocket.Upgrader{} // use default options

func food(w http.ResponseWriter, r *http.Request) {
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
		var msg string
		select {
		case food := <-channel:
			msg = food
		case <-time.After(30 * time.Second):
			msg = "no food"
		}
		err = c.WriteMessage(websocket.TextMessage, []byte(msg))
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
	http.HandleFunc("/food", food)
	go http.ListenAndServe(":8080", nil)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFoodCollectorServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
