package main

import (
	"context"
	"encoding/csv"
	pb "github.com/3hajk/grpc-project/proto"
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

func newServer() *routeGuideServer {
	s := &routeGuideServer{}
	return s
}

func saveData(data [][]string) error {
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
}
