package main

import (
	"context"
	"github.com/3hajk/grpc-project/proto"
	"log"
	"os"
	"testing"
	"time"
)

func TestFetch(t *testing.T) {
	ctx, logger, cancel := initTest("TestApproveAfterExpire", 60)
	defer cancel()

	grpc := newServer()

	request := &proto.FetchDataRequest{
		Url: "http://127.0.0.1:81/data.csv",
	}

	response, err := grpc.Fetch(ctx, request)
	if err != nil {
		logger.Fatal("", err.Error())
	}
	logger.Println("resp", response.String())

}

func TestList(t *testing.T) {
	ctx, logger, cancel := initTest("TestApproveAfterExpire", 60)
	defer cancel()

	grpc := newServer()

	page := &proto.Page{
		Skip:  0,
		Limit: 3,
	}
	sort := &proto.Sort{
		Name:   "name",
		Direct: 1,
	}
	request := &proto.ListProductsRequest{
		Paging:  page,
		Sorting: sort,
	}

	response, err := grpc.List(ctx, request)
	if err != nil {
		logger.Fatal("", err.Error())
	}
	logger.Println("resp", response.String())
	page = &proto.Page{
		Skip:  3,
		Limit: 3,
	}
	request = &proto.ListProductsRequest{
		Paging:  page,
		Sorting: sort,
	}
	response, err = grpc.List(ctx, request)
	if err != nil {
		logger.Fatal("", err.Error())
	}
	logger.Println("resp", response.String())

	page = &proto.Page{
		Skip:  6,
		Limit: 3,
	}
	request = &proto.ListProductsRequest{
		Paging:  page,
		Sorting: sort,
	}
	response, err = grpc.List(ctx, request)
	if err != nil {
		logger.Fatal("", err.Error())
	}
	logger.Println("resp", response.String())

	page = &proto.Page{
		Skip:  9,
		Limit: 3,
	}
	request = &proto.ListProductsRequest{
		Paging:  page,
		Sorting: sort,
	}
	response, err = grpc.List(ctx, request)
	if err != nil {
		logger.Fatal("", err.Error())
	}
	logger.Println("resp", response.String())

}

func initTest(testName string, testDuration int64) (ctx context.Context, logger *log.Logger, cancel func()) {
	logger = log.New(os.Stdout, testName, log.Ldate|log.Ltime|log.Lshortfile)
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(testDuration)*time.Second)

	return ctx, logger, cancel
}
