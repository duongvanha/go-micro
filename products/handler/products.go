package handler

import (
	"context"
	"github.com/samber/do"
	"io"
	"time"

	"go-micro.dev/v4/logger"

	pb "proto/services"
)

func NewProducts(injector *do.Injector) (*Products, error) {
	return &Products{}, nil
}

type Products struct{}

func (e *Products) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	logger.Infof("Received Products.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *Products) ClientStream(ctx context.Context, stream pb.Products_ClientStreamStream) error {
	var count int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			logger.Infof("Got %v pings total", count)
			return stream.SendMsg(&pb.ClientStreamResponse{Count: count})
		}
		if err != nil {
			return err
		}
		logger.Infof("Got ping %v", req.Stroke)
		count++
	}
}

func (e *Products) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.Products_ServerStreamStream) error {
	logger.Infof("Received Products.ServerStream request: %v", req)
	for i := 0; i < int(req.Count); i++ {
		logger.Infof("Sending %d", i)
		if err := stream.Send(&pb.ServerStreamResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 250)
	}
	return nil
}

func (e *Products) BidiStream(ctx context.Context, stream pb.Products_BidiStreamStream) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		logger.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.BidiStreamResponse{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
