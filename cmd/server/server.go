package main

import (
	"context"
	"log"
	"net"

	calc "github.com/zahaar/grcp-calculator/gen"
	"google.golang.org/grpc"
)

type calculatorService struct {
	calc.UnimplementedCalculatorServer
}

func (s *calculatorService) Add(_ context.Context, req *calc.AddRequest) (*calc.AddResponse, error) {
	log.Printf("Received %v; %v", req.GetA(), req.GetB())
	return &calc.AddResponse{S: req.GetA() + req.GetB()}, nil
}

func main() {
	// flag.Parse()
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	calc.RegisterCalculatorServer(grpcServer, &calculatorService{})
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
