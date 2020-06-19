package main

import (
	"context"
	"encoding/csv"
	"fmt"
	pb "github.com/3hajk/grpc-project/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
)

type routeGuideServer struct {
	pb.UnimplementedServerServer
	mu sync.Mutex // protects routeNotes
}

type row struct {
	Name    string
	Price   float64
	cont    int64
	updated uint64
}

func newServer() *routeGuideServer {
	s := &routeGuideServer{}
	return s
}

func saveData(data [][]string) error {

	fmt.Println(data)
	fmt.Println(data[0])
	fmt.Println(data[0][1])
	return nil
}

func (s *routeGuideServer) Fetch(ctx context.Context, req *pb.FetchRequest) (*pb.FetchResponse, error) {
	resp, err := http.Get(req.Url)
	if err != nil {
		return &pb.FetchResponse{Error: err.Error()}, err
	}
	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ';'

	data, err := reader.ReadAll()
	if err != nil {
		return &pb.FetchResponse{Error: err.Error()}, err
	}
	saveData(data)
	return &pb.FetchResponse{}, nil
}

func (s *routeGuideServer) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {

	return nil, nil
}

func main() {

	port := os.Getenv("PORT")

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	defer grpcServer.Stop()

	pb.RegisterServerServer(grpcServer, newServer())
	grpcServer.Serve(listener)

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Get a handle for your collection
	//collection := client.Database("test").Collection("trainers")

	//collection.FindOneAndUpdate()

}
