package main

import (
	"context"
	"errors"
	"log"
	"net"

	calc "github.com/zahaar/grcp-calculator/gen"
	"google.golang.org/grpc"
)

type calculatorService struct {
	calc.UnimplementedCalculatorServer
}

func calculate(r *calc.MathTaskRequest) (*calc.MathTaskResponse, error) {
	log.Printf("Received Method:{%v} Args:%v; %v", r.GetMethod(), r.GetArg1(), r.GetArg2())

	// check for zero division
	if r.GetArg1() == 0 || r.GetArg2() == 0 {
		return nil, errors.New("math: divided by zero")
	}

	var eq, arg1, arg2 float64 = 0, r.GetArg1(), r.GetArg2()
	switch r.GetMethod() {
	case calc.MathMethod_ADD:
		eq = arg1 + arg2
	case calc.MathMethod_SUB:
		eq = arg1 - arg2
	case calc.MathMethod_MUL:
		eq = arg1 * arg2
	case calc.MathMethod_DIV:
		eq = arg1 / arg2
	default:
		log.Fatal("Invalid Operation")
	}
	return &calc.MathTaskResponse{Eq: eq}, nil
}

func (s *calculatorService) PerformCalc(_ context.Context, req *calc.MathTaskRequest) (*calc.MathTaskResponse, error) {
	return calculate(req)
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("failed to listen:", err)
	}

	grpcServer := grpc.NewServer()
	calc.RegisterCalculatorServer(grpcServer, &calculatorService{})
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
