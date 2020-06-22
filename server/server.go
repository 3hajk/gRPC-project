package main

import (
	"context"
	"encoding/csv"
	"fmt"
	db "github.com/3hajk/grpc-project/database"
	pb "github.com/3hajk/grpc-project/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	pb.UnimplementedProductServiceServer
	mu sync.Mutex
	db *mongo.Client
}

func newServer() *routeGuideServer {
	s := &routeGuideServer{
		db: db.Connect(),
	}
	db.SetIndex(s.db)
	return s
}

func (s *routeGuideServer) Fetch(ctx context.Context, req *pb.FetchDataRequest) (*pb.FetchDataResponse, error) {
	resp, err := http.Get(req.Url)
	if err != nil {
		return &pb.FetchDataResponse{Error: err.Error()}, err
	}
	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ';'

	productList, err := reader.ReadAll()
	if err != nil {
		return &pb.FetchDataResponse{Error: err.Error()}, err
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
		err = db.InsertProductToDB(ctx, s.db, item)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return &pb.FetchDataResponse{}, nil
}

func (s *routeGuideServer) List(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	skip := req.GetPaging().GetSkip()
	limit := req.GetPaging().GetLimit()
	data, err := db.GetProductFromDB(ctx, s.db, skip, limit, req.GetSorting().GetName(), req.GetSorting().GetDirect())
	if err != nil {
		log.Println("Couldn't get sugar levels", "err", err)
		return nil, err
	}
	return &pb.ListProductsResponse{
		PageSize: int32(len(data)),
		List:     data,
	}, nil
}

func (s *routeGuideServer) Stream(req *pb.StreamProductsRequest, stream pb.ProductService_StreamServer) error {

	filter := bson.M{
		"Sort": bson.M{req.GetSorting().GetName(): req.GetSorting().GetDirect()},
	}
	data := &db.ProductItem{}

	collection := s.db.Database(db.DBNAME).Collection("products")
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		if err != nil {
			return err
		}
		stream.Send(&pb.StreamProductsResponse{
			Product: &pb.Product{
				Name:        data.Name,
				Price:       float32(data.Price),
				LastUpdate:  data.LastUpdate.T,
				ChangePrice: int32(data.Count),
			},
		})
		if err = cursor.Err(); err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("[-] Cursor Error: %v", err))
		}
	}
	return nil
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

	pb.RegisterProductServiceServer(grpcServer, newServer())

	// Start the server in a child routine
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	log.Println("Server succesfully started on port :", port)
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
