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

	request := &proto.FetchRequest{
		Url: "http://127.0.0.1:81/data2.csv",
	}

	response, err := grpc.Fetch(ctx, request)
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
