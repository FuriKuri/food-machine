//go:generate protoc -I ../food --go_out=plugins=grpc:../food ../food/food.proto

package main

import (
	"fmt"
	"net"
	"time"

	"net/http"

	pb "github.com/furikuri/food-machine/food"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"html/template"
	"log"
	"strconv"
	"sync/atomic"

	"github.com/gorilla/websocket"
)

const (
	port = ":50051"
)

type server struct{}

var listenerSockets = make(map[chan string]bool)
var ops uint64
var worker = make(map[string]bool)

func (s *server) Deliver(ctx context.Context, in *pb.Food) (*pb.Empty, error) {
	log.Print(in.Name)
	atomic.AddUint64(&ops, 1)
	worker[in.Worker] = true
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
		msg := <-channel
		err = c.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("index.html")).Execute(w, r.Host)
}

func updateStats() {
	ticker := time.NewTicker(1000 * time.Millisecond)
	for t := range ticker.C {
		fmt.Println("Update at", t)
		rate := atomic.SwapUint64(&ops, 0)
		workerCount := len(worker)
		worker = make(map[string]bool)
		for listener := range listenerSockets {
			listener <- "stats: " + strconv.FormatUint(rate, 10) + " food/second - " + strconv.Itoa(workerCount) + " worker"
		}
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/food", food)
	go http.ListenAndServe(":8080", nil)

	go updateStats()

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
