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

		//price, err := strconv.ParseFloat(row[1], 32)
		//if err != nil {
		//	price = 0
		//}
		item := db.ProductItem{
			Name:       row[0],
			Price:      row[1],
			LastUpdate: primitive.Timestamp{T: uint32(time.Now().Unix())},
			Count:      1,
		}
		fmt.Println(item)
		err = db.InsertProductToTheDB(ctx, s.db, item)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return &pb.FetchDataResponse{}, nil
}

//func (sr *routeGuideServer) GetBatchLevels(from, to, skip, limit int64, userId string) ([]*dto.SugarLevel, error) {
//	ctx := context.Background()
//	collection := sr.db.Database(sr.GetDbName()).Collection(srName)
//	opts := &options.FindOptions{
//		Skip:&skip,
//		Limit:&limit,
//		Sort:bson.D{
//			{"timestamp", 1}}}
//	cursor, err := collection.Find(ctx, bson.D{
//		{  "userid",userId},
//		{  "timestamp", bson.D{
//			{"$gte", from}, {"$lte", to}},
//		}}, opts)
//	if err != nil {  log.Println("Couldn't get sugar levels", "err", err)  return nil, err }
//	defer cursor.Close(ctx)
//	results := make([]*dto.SugarLevel, 0, 0)
//	for cursor.Next(context.Background()) {
//		var result = new(dto.SugarLevel)
//		err := cursor.Decode(result)
//		if err != nil {   return results, err  }
//		results = append(results, result)
//	}
//	return results, nil
//}

func (s *routeGuideServer) List(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {

	return nil, nil
}

func (s *routeGuideServer) Stream(req *pb.StreamProductsRequest, stream ProductService_StreamServer) error {

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
