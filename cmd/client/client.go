package main

import (
	"context"
	"flag"
	"log"
	"time"

	calc "github.com/zahaar/grcp-calculator/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	flag.Parse()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial("localhost:8080", opts...)

	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	c := calc.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.Add(ctx, &calc.AddRequest{A: 21, B: 21})
	if err != nil {
		log.Fatalf("fail to perform Add()")
	}
	log.Printf("Result: %v", res.GetS())
}
