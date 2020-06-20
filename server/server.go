package main

import (
	"context"
	"encoding/csv"
	"fmt"
	db "github.com/3hajk/grpc-project/database"
	pb "github.com/3hajk/grpc-project/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

type routeGuideServer struct {
	pb.UnimplementedServerServer
	mu sync.Mutex
	db *mongo.Client
}

func newServer() *routeGuideServer {
	s := &routeGuideServer{
		db: db.Connect(),
	}
	return s
}

func (s *routeGuideServer) Fetch(ctx context.Context, req *pb.FetchRequest) (*pb.FetchResponse, error) {
	resp, err := http.Get(req.Url)
	if err != nil {
		return &pb.FetchResponse{Error: err.Error()}, err
	}
	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ';'

	productList, err := reader.ReadAll()
	if err != nil {
		return &pb.FetchResponse{Error: err.Error()}, err
	}
	for _, row := range productList {
		fmt.Println(row)

		price, err := strconv.ParseFloat(row[1], 32)
		if err != nil {
			price = 0
		}
		item := db.ProductItem{
			Name:       row[0],
			Price:      price,
			LastUpdate: primitive.Timestamp{T: uint32(time.Now().Unix())},
			Count:      1,
		}
		fmt.Println(item)
		_ = db.InsertProductToTheDB(ctx, s.db, item)
	}
	return &pb.FetchResponse{}, nil
}

func (s *routeGuideServer) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {

	return nil, nil
}

func main() {

	port := os.Getenv("PORT")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	defer grpcServer.Stop()

	pb.RegisterServerServer(grpcServer, newServer())

	// Start the server in a child routine
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("Server succesfully started on port :", port)
	// Create a channel to receive OS signals
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c
	// After receiving CTRL+C Properly stop the server
	log.Println("\nStopping the server...")
	grpcServer.Stop()
	listener.Close()
	log.Println("Done.")
}
